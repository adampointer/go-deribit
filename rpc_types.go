package deribit

import (
	"bytes"
	"encoding/json"
	"github.com/go-openapi/runtime"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const rpcVersion = "2.0"

// RPCRequest is what we send to the remote
// Implements runtime.ClientRequest
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

func (RPCRequest) SetHeaderParam(string, ...string) error {
	return nil
}

func (RPCRequest) GetHeaderParams() http.Header {
	return http.Header{}
}

func (r *RPCRequest) SetQueryParam(key string, vals ...string) error {
	if len(vals) > 1 {
		r.Params[key] = vals
	} else {
		r.Params[key] = vals[0]
	}

	return nil
}

func (RPCRequest) SetFormParam(string, ...string) error {
	return nil
}

func (RPCRequest) SetPathParam(string, string) error {
	return nil
}

func (RPCRequest) GetQueryParams() url.Values {
	return url.Values{}
}

func (RPCRequest) SetFileParam(string, ...runtime.NamedReadCloser) error {
	return nil
}

func (RPCRequest) SetBodyParam(interface{}) error {
	return nil
}

func (RPCRequest) SetTimeout(time.Duration) error {
	return nil
}

func (RPCRequest) GetMethod() string {
	return "GET"
}

func (r *RPCRequest) GetPath() string {
	return r.Method
}

func (RPCRequest) GetBody() []byte {
	return []byte{}
}

func (RPCRequest) GetBodyParam() interface{} {
	return nil
}

func (RPCRequest) GetFileParam() map[string][]runtime.NamedReadCloser {
	return nil
}

// RPCResponse is what we receive from the remote
// Implements runtime.ClientResponse
type RPCResponse struct {
	JsonRpc string          `json:"jsonrpc"`
	ID      uint64          `json:"id,omitempty"`
	Result  json.RawMessage `json:"result"`
	Error   *RPCError       `json:"error,omitempty"`
}

func (r *RPCResponse) Code() int {
	if r.Error != nil {
		return r.Error.Code
	}
	return 200
}

func (r *RPCResponse) Message() string {
	if r.Error != nil {
		return r.Error.Message
	}
	return "ok"
}

func (RPCResponse) GetHeader(string) string {
	return ""
}

func (r *RPCResponse) Body() io.ReadCloser {
	b, _ := json.Marshal(r)
	return ioutil.NopCloser(bytes.NewReader(b))
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
		Data    json.RawMessage `json:"data"`
		Channel string          `json:"channel"`
	} `json:"params,omitempty"`
}
