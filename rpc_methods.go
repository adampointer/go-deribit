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
func (e *Exchange) GetLastTrades(count, since int) (*RPCResponse, []*TradeResponse, error) {
	req := RPCRequest{Action: "/api/v1/public/getlasttrades"}
	req.Arguments = map[string]interface{}{"instrument": "BTC-PERPETUAL"}
	if count != 0 {
		req.Arguments["count"] = count
	}
	if since != 0 {
		req.Arguments["since"] = since
	}
	res, err := e.makeRequest(req)
	if err != nil {
		return nil, nil, err
	}
	trades := make([]*TradeResponse, 0)
	if len(res.Result) != 0 {
		if err := json.Unmarshal(res.Result, &trades); err != nil {
			return nil, nil, fmt.Errorf("Unable to unmarshal result: %s", err)
		}
	}
	return res, trades, nil
}

// Buy places a buy order
func (e *Exchange) Buy(or *OrderRequest) (*RPCResponse, *OrderResponse, error) {
	return e.placeOrder("buy", or)
}

// Sell places a buy order
func (e *Exchange) Sell(or *OrderRequest) (*RPCResponse, *OrderResponse, error) {
	return e.placeOrder("buy", or)
}

func (e *Exchange) placeOrder(action string, or *OrderRequest) (*RPCResponse, *OrderResponse, error) {
	if len(action) == 0 {
		return nil, nil, fmt.Errorf("Action must be passed")
	}
	args, err := or.toMap()
	if err != nil {
		return nil, nil, err
	}
	req := RPCRequest{
		Action:    "/api/v1/private/" + action,
		Arguments: args,
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
