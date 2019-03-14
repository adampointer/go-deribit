package deribit

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
<<<<<<< Updated upstream
	"sort"
	"strconv"
=======
	"github.com/mitchellh/mapstructure"
>>>>>>> Stashed changes
	"strings"
	"time"
)

// RPCRequest is what we send to the remote
type RPCRequest struct {
	Action    string                 `json:"action"`
	ID        uint64                 `json:"id"`
	Arguments map[string]interface{} `json:"arguments,omitempty"`
	Sig       string                 `json:"sig,omitempty"`
}

<<<<<<< Updated upstream
// GenerateSig creates the signature required for private endpoints
func (r *RPCRequest) GenerateSig(key, secret string) error {
	if len(key) == 0 || len(secret) == 0 {
		return errors.New("You must supply an access key and an access secret")
	}
	nonce := time.Now().UnixNano() / int64(time.Millisecond)
	sigString := fmt.Sprintf("_=%d&_ackey=%s&_acsec=%s&_action=%s", nonce, key, secret, r.Action)

	// Append args if present
	if len(r.Arguments) != 0 {
		var argsString string

		// We have to do this to sort by keys
		keys := make([]string, 0, len(r.Arguments))
		for key := range r.Arguments {
			keys = append(keys, key)
		}
		sort.Strings(keys)

		for _, k := range keys {
			v := r.Arguments[k]
			var s string

			switch t := v.(type) {
			case []SubscriptionEvent:
				var str = make([]string, len(t))
				for _, j := range t {
					str = append(str, string(j))
				}
				s = strings.Join(str, "")
			case []string:
				s = strings.Join(t, "")
			case bool:
				s = strconv.FormatBool(t)
			case int:
				s = strconv.FormatInt(int64(t), 10)
			case int64:
				s = strconv.FormatInt(t, 10)
			case float64:
				s = strconv.FormatFloat(t, 'f', -1, 64)
			case string:
				s = t
			default:
				// Absolutely panic here
				panic(fmt.Sprintf("Cannot generate sig string: Unable to handle arg of type %T", t))
			}
			argsString += fmt.Sprintf("&%s=%s", k, s)
		}
		sigString += argsString
	}
	hasher := sha256.New()
	hasher.Write([]byte(sigString))
	sigHash := base64.StdEncoding.EncodeToString(hasher.Sum(nil))
	r.Sig = fmt.Sprintf("%s.%d.%s", key, nonce, sigHash)
	return nil
=======
func NewRPCRequest(method string) *RPCRequest {
	return &RPCRequest{
		JsonRpc: rpcVersion,
		Method:  method,
		Params:  make(map[string]interface{}),
	}
>>>>>>> Stashed changes
}

// RPCResponse is what we receive from the remote
type RPCResponse struct {
	ID            uint64             `json:"id"`
	Success       bool               `json:"success"`
	Error         int                `json:"error"`
	Testnet       bool               `json:"testnet"`
	Message       string             `json:"message"`
	UsIn          uint64             `json:"usIn"`
	UsOut         uint64             `json:"usOut"`
	UsDiff        uint64             `json:"usDiff"`
	Result        json.RawMessage    `json:"result"`
	APIBuild      string             `json:"apiBuild,omitempty"`
	Notifications []*RPCNotification `json:"notifications,omitempty"`
}

// RPCCall represents the entire call from request to response
type RPCCall struct {
	Req   RPCRequest
	Res   RPCResponse
	Error error
	Done  chan bool
}

// NewRPCCall returns a new RPCCall initialised with a done channel and request
func NewRPCCall(req RPCRequest) *RPCCall {
	done := make(chan bool)
	return &RPCCall{
		Req:  req,
		Done: done,
	}
}

<<<<<<< Updated upstream
// makeRequest makes a request over the websocket and waits for a response with a timeout
func (e *Exchange) makeRequest(req RPCRequest) (*RPCResponse, error) {
=======
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
>>>>>>> Stashed changes
	e.mutex.Lock()
	id := e.counter
	e.counter++
	req.ID = id
	call := NewRPCCall(req)
	e.pending[id] = call

	if err := e.conn.WriteJSON(&req); err != nil {
		delete(e.pending, id)
		e.mutex.Unlock()
		return nil, err
	}
	e.mutex.Unlock()
	select {
	case <-call.Done:
	case <-time.After(10 * time.Second):
		call.Error = fmt.Errorf("Request %d timed out", id)
	}

	if call.Error != nil {
		return nil, call.Error
	}
	if !call.Res.Success {
		return nil, fmt.Errorf("Request failed with code (%d): %s", call.Res.Error, call.Res.Message)
	}
	return &call.Res, nil
}

// read takes messages off the websocket and deals with them accordingly
// This method is ugly AF as needs refactoring
func (e *Exchange) read() {
	var resErr error
Loop:
	for {
		select {
		case <-e.stop:
			break Loop
		default:
			var res RPCResponse
			if err := e.conn.ReadJSON(&res); err != nil {
				e.errors <- fmt.Errorf("Error reading message: %q", err)
			}
<<<<<<< Updated upstream

			// Notifications do not have an ID field
			if res.Notifications != nil {
				for _, n := range res.Notifications {
					e.mutex.Lock()
					sub := e.subscriptions[n.Message]
					e.mutex.Unlock()
					if sub == nil {
						// Send error to main error channel
						e.errors <- fmt.Errorf("No subscription found for %s", n.Message)
					}
					// Send the notification to the right channel
					sub.Data <- n
=======

			f := func(key string) bool { _, ok := raw[key]; return ok }

			if f("id") || f("error") {
				var res RPCResponse
				if err := mapstructure.Decode(raw, &res); err != nil {
					resErr = fmt.Errorf("unable to to decode message to RPCResponse: %s", err)
					break Loop
>>>>>>> Stashed changes
				}
			} else {
				e.mutex.Lock()
				call := e.pending[res.ID]
				e.mutex.Unlock()

<<<<<<< Updated upstream
				if call == nil {
					resErr = fmt.Errorf("No pending request found for response ID %d", res.ID)
=======
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
			} else if _, ok := raw["method"]; ok {
				var res RPCNotification
				if err := mapstructure.Decode(raw, &res); err != nil {
					resErr = fmt.Errorf("unable to to decode message to RPCNotification: %s", err)
>>>>>>> Stashed changes
					break Loop
				}
				call.Res = res
				call.Done <- true
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
