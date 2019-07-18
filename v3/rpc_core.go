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

type RPCCore struct {
	conn          *websocket.Conn
	reqMutex      sync.Mutex
	pending       map[uint64]*RPCCall
	subscriptions map[string]*RPCSubscription
	counter       uint64
	closedMutex   sync.Mutex
	isClosed      bool
	onDisconnect  func(*RPCCore)
	errors        chan error
	stop          chan bool
	auth          *models.PublicAuthResponse
}

// Submit satisfies the runtime.ClientTransport interface
func (r *RPCCore) Submit(operation *runtime.ClientOperation) (interface{}, error) {
	method := operation.PathPattern
	// Strip leading slashes
	method = strings.TrimPrefix(method, "/")
	req := NewRPCRequest(method)
	if err := operation.Params.WriteToRequest(req, strfmt.Default); err != nil {
		return nil, err
	}
	// Add auth
	if strings.HasPrefix(method, "private/") && r.auth != nil {
		req.Params["access_token"] = r.auth.Result.AccessToken
	}
	res, err := r.rpcRequest(req)
	if err != nil {
		return nil, err
	}
	return operation.Reader.ReadResponse(res, runtime.JSONConsumer())
}

func (r *RPCCore) close() error {
	r.closedMutex.Lock()
	r.isClosed = true
	r.closedMutex.Unlock()

	if err := r.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")); err != nil {
		return err
	}
	return r.conn.Close()
}

func (r *RPCCore) rpcRequest(req *RPCRequest) (*RPCResponse, error) {
	call := NewRPCCall(req)
	// Create a new request ID
	r.reqMutex.Lock()
	id := r.counter
	r.counter++
	req.ID = id
	r.pending[id] = call

	// Send
	if err := r.conn.WriteJSON(&req); err != nil {
		delete(r.pending, id)
		r.reqMutex.Unlock()
		return nil, err
	}
	r.reqMutex.Unlock()

	// Wait for response or timeout
	select {
	case <-call.Done:
	case <-time.After(3 * time.Second):
		call.Error = fmt.Errorf("request %d timed out", id)
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
			if err := r.conn.ReadJSON(&raw); err != nil {
				r.closedMutex.Lock()
				isClosed := r.isClosed
				r.closedMutex.Unlock()

				// fix for `use of closed network connection`
				if isClosed {
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
		r.reqMutex.Lock()
		for _, call := range r.pending {
			call.Error = resErr
			call.Done <- true
		}
		r.reqMutex.Unlock()
	}
}

func (r *RPCCore) handleResponses(res *RPCResponse) error {
	r.reqMutex.Lock()
	call := r.pending[res.ID]
	r.reqMutex.Unlock()

	if strings.Contains(call.Req.Method, "subscribe") {
		if len(res.Result) <= 2 && res.Error == nil {
			res.Error = &RPCError{Code: 10001, Message: "empty result"}
		}
	}

	if res.Error != nil && res.Error.Code != 0 {
		resErr := fmt.Errorf("request failed with code (%d): %s", res.Error.Code, res.Error.Message)
		call.Error = resErr
		call.Done <- true
	} else {
		if call == nil {
			return fmt.Errorf("no pending request found for response ID %d", res.ID)
		}
		call.Res = res
		call.Done <- true
		r.reqMutex.Lock()
		delete(r.pending, res.ID)
		r.reqMutex.Unlock()
	}
	return nil
}

func (r *RPCCore) handleSubscriptions(res *RPCNotification) {
	r.reqMutex.Lock()
	sub := r.subscriptions[res.Params.Channel]
	r.reqMutex.Unlock()
	if sub == nil {
		// Send error to main error channel
		r.errors <- fmt.Errorf("no subscription found for %s", res.Params.Channel)
		return
	}
	// Send the notification to the right channel
	sub.Data <- res
}
