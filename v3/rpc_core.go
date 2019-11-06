package deribit

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/adampointer/go-deribit/v3/models"

	"github.com/gorilla/websocket"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

type callManager struct {
	reqMutex      sync.Mutex
	pending       map[uint64]*RPCCall
	subMutex      sync.Mutex
	subscriptions map[string]*RPCSubscription
	countMutex    sync.Mutex
	counter       uint64
}

func (c *callManager) nextID() uint64 {
	c.countMutex.Lock()
	defer c.countMutex.Unlock()
	id := c.counter
	c.counter++
	return id
}

func (c *callManager) addCall(call *RPCCall) {
	// Create a new request ID
	c.reqMutex.Lock()
	call.Req.ID = c.nextID()
	c.pending[call.Req.ID] = call
	c.reqMutex.Unlock()
}

func (c *callManager) deleteCall(call *RPCCall) {
	c.reqMutex.Lock()
	delete(c.pending, call.Req.ID)
	c.reqMutex.Unlock()
}

func (c *callManager) getCall(id uint64) *RPCCall {
	c.reqMutex.Lock()
	defer c.reqMutex.Unlock()
	return c.pending[id]
}

func (c *callManager) closeAllWithError(err error) {
	c.reqMutex.Lock()
	for _, call := range c.pending {
		call.Error = err
		call.Done <- true
	}
	c.reqMutex.Unlock()
}

func (c *callManager) addSubscription(sub *RPCSubscription) {
	c.subMutex.Lock()
	c.subscriptions[sub.Channel] = sub
	c.subMutex.Unlock()
}

func (c *callManager) deleteSubscription(channel string) {
	c.subMutex.Lock()
	delete(c.subscriptions, channel)
	c.subMutex.Unlock()
}

func (c *callManager) getSubscription(channel string) *RPCSubscription {
	c.subMutex.Lock()
	defer c.subMutex.Unlock()
	return c.subscriptions[channel]
}

func (c *callManager) getSubscriptions() map[string]*RPCSubscription {
	c.subMutex.Lock()
	defer c.subMutex.Unlock()
	return c.subscriptions
}

type connManager struct {
	stopPinging func()
	closedMutex sync.Mutex
	isClosed    bool
	writeMutex  sync.Mutex
	conn        *websocket.Conn
}

func (c *connManager) closed() bool {
	c.closedMutex.Lock()
	defer c.closedMutex.Unlock()
	return c.isClosed
}

func (c *connManager) close() error {
	c.closedMutex.Lock()
	c.isClosed = true
	c.closedMutex.Unlock()
	if err := c.writeMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")); err != nil {
		return err
	}
	return c.conn.Close()
}

func (c *connManager) writeJSON(msg interface{}) error {
	c.writeMutex.Lock()
	defer c.writeMutex.Unlock()
	return c.conn.WriteJSON(msg)
}

func (c *connManager) writeMessage(msgType int, msg []byte) error {
	c.writeMutex.Lock()
	defer c.writeMutex.Unlock()
	return c.conn.WriteMessage(msgType, msg)
}

func (c *connManager) readJSON(msg interface{}) error {
	return c.conn.ReadJSON(msg)
}

// RPCCore actually sends and receives messages
type RPCCore struct {
	OnConnect func()

	calls        *callManager
	connMgr      *connManager
	onDisconnect func(*RPCCore)
	errors       chan error
	stop         chan bool
	auth         *models.PublicAuthResponse
}

// Submit satisfies the runtime.ClientTransport interface
func (r *RPCCore) Submit(operation *runtime.ClientOperation) (interface{}, error) {
	// Strip leading slashes
	method := strings.TrimPrefix(operation.PathPattern, "/")
	req := NewRPCRequest(method)
	if err := operation.Params.WriteToRequest(req, strfmt.Default); err != nil {
		return nil, err
	}
	req.AddAuth(r.auth)
	res, err := r.rpcRequest(req)
	if err != nil {
		return nil, err
	}
	return operation.Reader.ReadResponse(res, runtime.JSONConsumer())
}

func (r *RPCCore) close() error {
	return r.connMgr.close()
}

func (r *RPCCore) rpcRequest(req *RPCRequest) (*RPCResponse, error) {
	call := NewRPCCall(req)
	r.calls.addCall(call)
	// Send
	r.connMgr.conn.SetWriteDeadline(time.Now().Add(writeWait))
	if err := r.connMgr.writeJSON(&req); err != nil {
		r.calls.deleteCall(call)
		return nil, err
	}

	// Wait for response or timeout
	select {
	case <-call.Done:
	case <-time.After(3 * time.Second):
		call.Error = fmt.Errorf("request %d timed out", call.Req.ID)
	}
	if call.Error != nil {
		return nil, call.Error
	}
	if call.Res.Error != nil {
		return nil, fmt.Errorf("request failed with code (%d): %s", call.Res.Error.Code, call.Res.Error.Message)
	}
	return call.Res, nil
}

// read takes messages off the websocket and deals with them accordingly
func (r *RPCCore) read() {
	var resErr error
Loop:
	for {
		select {
		case <-r.stop:
			break Loop
		default:
			var raw composite
			r.connMgr.conn.SetReadDeadline(time.Now().Add(pongWait))
			if err := r.connMgr.readJSON(&raw); err != nil {
				// fix for `use of closed network connection`
				if r.connMgr.closed() {
					break Loop
				}
				r.onDisconnect(r)
				break Loop
			}

			if raw.ID != 0 || raw.Error != nil {
				resErr = r.handleResponses(raw.toResponse())
			} else if raw.Method == "subscription" {
				r.handleSubscriptions(raw.toNotification())
			}
		}
	}
	if resErr != nil {
		r.calls.closeAllWithError(resErr)
	}
}

func (r *RPCCore) handleResponses(res *RPCResponse) error {
	call := r.calls.getCall(res.ID)

	if strings.Contains(call.Req.Method, "subscribe") {
		if len(res.Result) <= 2 && res.Error == nil {
			res.Error = &RPCError{Code: 10001, Message: "empty result"}
		}
	}

	if res.Error != nil && res.Error.Code != 0 {
		resErr := fmt.Errorf("request failed with code (%d): %s", res.Error.Code, res.Error.Message)
		call.CloseError(resErr)
	} else {
		if call == nil {
			return fmt.Errorf("no pending request found for response ID %d", res.ID)
		}
		call.Res = res
		call.CloseOK()
		r.calls.deleteCall(call)
	}
	return nil
}

func (r *RPCCore) handleSubscriptions(res *RPCNotification) {
	sub := r.calls.getSubscription(res.Params.Channel)
	if sub == nil {
		// Send error to main error channel
		r.errors <- fmt.Errorf("no subscription found for %s", res.Params.Channel)
		return
	}
	// Send the notification to the right channel
	sub.Data <- res
}
