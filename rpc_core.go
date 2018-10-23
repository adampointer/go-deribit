package deribit

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
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

// GenerateSig creates the signature required for private endpoints
func (r *RPCRequest) GenerateSig(key, secret string) error {
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

// makeRequest makes a request over the websocket and waits for a response with a timeout
func (e *Exchange) makeRequest(req RPCRequest) (*RPCResponse, error) {
	e.mutex.Lock()
	id := e.counter
	e.counter++
	req.ID = id
	call := NewRPCCall(req)
	e.pending[id] = call

	//j, _ := json.Marshal(req)
	//fmt.Printf("Req: %s\n", j)

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

			//j, _ := json.Marshal(res)
			//fmt.Printf("Res: %s\n", j)

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
				}
			} else {
				e.mutex.Lock()
				call := e.pending[res.ID]
				delete(e.pending, res.ID)
				e.mutex.Unlock()

				if call == nil {
					resErr = fmt.Errorf("No pending request found for response ID %d", res.ID)
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
