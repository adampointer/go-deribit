package deribit

import (
	"fmt"
	"strings"
	"time"

	"github.com/adampointer/go-deribit/v3/client/operations"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

type composite struct {
	RPCNotification
	RPCResponse
}

// Submit satisfies the runtime.ClientTransport interface
func (e *Exchange) Submit(operation *runtime.ClientOperation) (interface{}, error) {
	method := operation.PathPattern
	// Strip leading slashes
	method = strings.TrimPrefix(method, "/")
	req := NewRPCRequest(method)
	if err := operation.Params.WriteToRequest(req, strfmt.Default); err != nil {
		return nil, err
	}
	// Add auth
	if strings.HasPrefix(method, "private/") && e.auth != nil {
		req.Params["access_token"] = e.auth.Result.AccessToken
	}
	res, err := e.rpcRequest(req)
	if err != nil {
		return nil, err
	}
	return operation.Reader.ReadResponse(res, runtime.JSONConsumer())
}

// Client returns an initialised API client
func (e *Exchange) Client() *operations.Client {
	if e.client == nil {
		e.client = operations.New(e, strfmt.Default)
	}
	return e.client
}

func (e *Exchange) rpcRequest(req *RPCRequest) (*RPCResponse, error) {
	call := NewRPCCall(req)
	// Create a new request ID
	e.mutex.Lock()
	id := e.counter
	e.counter++
	req.ID = id
	e.pending[id] = call

	// Send
	if err := e.conn.WriteJSON(&req); err != nil {
		delete(e.pending, id)
		e.mutex.Unlock()
		return nil, err
	}
	e.mutex.Unlock()

	// Wait for response or timeout
	select {
	case <-call.Done:
	case <-time.After(10 * time.Second):
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
func (e *Exchange) read() {
	var resErr error
Loop:
	for {
		select {
		case <-e.stop:
			break Loop
		default:
			var raw composite
			if err := e.conn.ReadJSON(&raw); err != nil {
				if isTemporary(err) {
					continue
				}
				// stop reading if a close message sent from server
				if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
					break Loop
				}
				if f := e.OnDisconnect; f != nil { // reconnect
					f(e)
				}
				break Loop
			}

			if raw.ID != 0 || raw.Error != nil {
				res := &RPCResponse{
					JsonRpc: rpcVersion,
					ID:      raw.ID,
					Result:  raw.Result,
					Error:   raw.Error,
				}
				if len(res.Result) <= 2 && res.Error == nil {
					res.Error = &RPCError{Code: 10001, Message: "empty result"}
				}
				e.mutex.Lock()
				call := e.pending[res.ID]
				e.mutex.Unlock()

				if res.Error != nil && res.Error.Code != 0 {
					resErr = fmt.Errorf("request failed with code (%d): %s", res.Error.Code, res.Error.Message)
					break Loop
				} else {
					if call == nil {
						resErr = fmt.Errorf("no pending request found for response ID %d", res.ID)
						break Loop
					}
					call.Res = res
					call.Done <- true
					e.mutex.Lock()
					delete(e.pending, res.ID)
					e.mutex.Unlock()
				}
			} else if raw.Method == "subscription" {
				res := &RPCNotification{
					JsonRpc: rpcVersion,
					Method:  raw.Method,
					Params:  raw.Params,
				}
				e.mutex.Lock()
				sub := e.subscriptions[res.Params.Channel]
				e.mutex.Unlock()
				if sub == nil {
					// Send error to main error channel
					e.errors <- fmt.Errorf("no subscription found for %s", res.Params.Channel)
				}
				// Send the notification to the right channel
				sub.Data <- res
			}
		}
	}
	if resErr != nil {
		e.mutex.Lock()
		for _, call := range e.pending {
			call.Error = resErr
			call.Done <- true
		}
		e.mutex.Unlock()
	}
}

type temporary interface {
	Temporary() bool
}

// returns true if network err is temporary.
func isTemporary(err error) bool {
	te, ok := errors.Cause(err).(temporary)
	return ok && te.Temporary()
}
