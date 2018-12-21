package deribit

import (
	"encoding/json"
	"fmt"
)

// Test makes a request to the public/test endpoint
func (e *Exchange) Test(forceException bool) (*RPCResponse, error) {
	req := RPCRequest{Action: "/api/v1/public/test"}
	if forceException {
		req.Arguments = map[string]interface{}{"expired": "true"}
	}
	return e.makeRequest(req)
}

// Ping the remote to check for stale connections
func (e *Exchange) Ping() (*RPCResponse, error) {
	req := RPCRequest{Action: "/api/v1/public/ping"}
	return e.makeRequest(req)
}

// GetLastTrades returns the latest trades for an instrument
func (e *Exchange) GetLastTrades(count int, startSeq, endSeq int64, instrument string) (*RPCResponse, []*TradeResponse, error) {
	req := RPCRequest{Action: "/api/v1/public/getlasttrades"}
	req.Arguments = map[string]interface{}{"instrument": instrument}
	if count != 0 {
		req.Arguments["count"] = count
	}
	if startSeq != 0 {
		req.Arguments["startSeq"] = startSeq
	}
	if endSeq != 0 {
		req.Arguments["endSeq"] = startSeq
	}
	res, err := e.makeRequest(req)
	if err != nil {
		return nil, nil, err
	}
	trades := make([]*TradeResponse, 0)
	if len(res.Result) != 0 {
		if err := json.Unmarshal(res.Result, &trades); err != nil {
			return nil, nil, fmt.Errorf("unable to unmarshal result: %s", err)
		}
	}
	return res, trades, nil
}

// Buy places a buy order
func (e *Exchange) Buy(or *OrderRequest) (*RPCResponse, *OrderResponse, error) {
	return e.placeOrder("buy", or)
}

// Sell places a sell order
func (e *Exchange) Sell(or *OrderRequest) (*RPCResponse, *OrderResponse, error) {
	return e.placeOrder("sell", or)
}

// Edit edits an existing order
func (e *Exchange) Edit(oid string, qty int, price, stopPrice float64, postOnly bool) (*RPCResponse, *OrderResponse, error) {
	req := RPCRequest{
		Action: "/api/v1/private/edit",
		Arguments: map[string]interface{}{
			"orderId":   oid,
			"quantity":  qty,
			"post_only": postOnly,
		},
	}
	if stopPrice > 0 {
		req.Arguments["stopPx"] = stopPrice
	}
	if price > 0 {
		req.Arguments["price"] = price
	}
	if err := req.GenerateSig(e.key, e.secret); err != nil {
		return nil, nil, err
	}
	res, err := e.makeRequest(req)
	if err != nil {
		return nil, nil, err
	}
	var ret OrderResponse
	if err := json.Unmarshal(res.Result, &ret); err != nil {
		return nil, nil, fmt.Errorf("unable to unmarshal result: %s", err)
	}
	return res, &ret, nil
}

// CancelAll orders
func (e *Exchange) CancelAll(instrument string) (*RPCResponse, error) {
	req := RPCRequest{
		Action: "/api/v1/private/cancelall",
		Arguments: map[string]interface{}{
			"instrument": instrument,
			"type":       "futures",
		},
	}
	if err := req.GenerateSig(e.key, e.secret); err != nil {
		return nil, err
	}
	return e.makeRequest(req)
}

func (e *Exchange) placeOrder(action string, or *OrderRequest) (*RPCResponse, *OrderResponse, error) {
	if len(action) == 0 {
		return nil, nil, fmt.Errorf("Action must be passed")
	}
	req := RPCRequest{
		Action:    "/api/v1/private/" + action,
		Arguments: or.toMap(),
	}
	if err := req.GenerateSig(e.key, e.secret); err != nil {
		return nil, nil, err
	}
	res, err := e.makeRequest(req)
	if err != nil {
		return nil, nil, err
	}
	var ret OrderResponse
	if err := json.Unmarshal(res.Result, &ret); err != nil {
		return nil, nil, fmt.Errorf("Unable to unmarshal result: %s", err)
	}
	return res, &ret, nil
}
