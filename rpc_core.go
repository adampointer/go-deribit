package deribit

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"strings"
	"time"
)

const rpcVersion = "2.0"

// RPCRequest is what we send to the remote
type RPCRequest struct {
	JsonRpc string                 `json:"jsonrpc"`
	Method  string                 `json:"method"`
	ID      uint64                 `json:"id"`
	Params  map[string]interface{} `json:"params,omitempty"`
}

func NewRPCRequest(method string) *RPCRequest {
	return &RPCRequest{
		JsonRpc: rpcVersion,
		Method:  method,
		Params:  make(map[string]interface{}),
	}
}

// RPCResponse is what we receive from the remote
type RPCResponse struct {
	JsonRpc string                 `json:"jsonrpc"`
	ID      uint64                 `json:"id,omitempty"`
	Result  map[string]interface{} `json:"result"`
	Error   *RPCError              `json:"error,omitempty"`
}

// RPCError error object
type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// RPCCall represents the entire call from request to response
type RPCCall struct {
	Req   *RPCRequest
	Res   *RPCResponse
	Error error
	Done  chan bool
}

// NewRPCCall returns a new RPCCall initialised with a done channel and request
func NewRPCCall(req *RPCRequest) *RPCCall {
	done := make(chan bool)
	return &RPCCall{
		Req:  req,
		Done: done,
	}
}

// RPCSubscription is a subscription to an event type to receive notifications about
type RPCSubscription struct {
	Data    chan *RPCNotification
	Channel string
}

// RPCNotification is a notification which we have subscribed to
type RPCNotification struct {
	JsonRpc string `json:"jsonrpc"`
	Method  string `json:"action"`
	Params  struct {
		Data    map[string]interface{} `json:"data"`
		Channel string                 `json:"channel"`
	} `json:"params,omitempty"`
}

func (e *Exchange) rpcRequest(method string, params interface{}) (*RPCResponse, error) {
	req := NewRPCRequest(method)
	call := NewRPCCall(req)

	// Encode parameters to map
	if params != nil {
		if err := mapstructure.Decode(params, &req.Params); err != nil {
			return nil, fmt.Errorf("error encoding parameters: %s", err)
		}
	}
	if strings.HasPrefix(method, "private/") && e.auth != nil {
		req.Params["access_token"] = e.auth.AccessToken
	}
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
			var raw map[string]interface{}
			if err := e.conn.ReadJSON(&raw); err != nil {
				e.errors <- fmt.Errorf("error reading message: %q", err)
			}

			f := func(key string) bool { _, ok := raw[key]; return ok }

			if f("id") || f("error") {
				var res RPCResponse
				// Type munging :)
				switch t := raw["result"].(type) {
				case []interface{}:
					result := make(map[string]interface{})
					for i, v := range t {
						result[fmt.Sprintf("%d", i)] = v
					}
					raw["result"] = result
				}
				if err := mapstructure.Decode(raw, &res); err != nil {
					resErr = fmt.Errorf("unable to to decode message to RPCResponse: %s", err)
					break Loop
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
					call.Res = &res
					call.Done <- true
					e.mutex.Lock()
					delete(e.pending, res.ID)
					e.mutex.Unlock()
				}
			} else if f("method") && raw["method"] == "subscription" {
				var res RPCNotification
				if err := mapstructure.Decode(raw, &res); err != nil {
					resErr = fmt.Errorf("unable to to decode message to RPCNotification: %s", err)
					break Loop
				}
				e.mutex.Lock()
				sub := e.subscriptions[res.Params.Channel]
				e.mutex.Unlock()
				if sub == nil {
					// Send error to main error channel
					e.errors <- fmt.Errorf("no subscription found for %s", res.Params.Channel)
				}
				// Send the notification to the right channel
				sub.Data <- &res
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
