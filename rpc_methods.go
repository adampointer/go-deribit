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
	return e.makeRequest(req, "")
}

// Ping the remote to check for stale connections
func (e *Exchange) Ping() (*RPCResponse, error) {
	req := RPCRequest{Action: "/api/v1/public/ping"}
	return e.makeRequest(req, "")
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
	res, err := e.makeRequest(req, "trade")
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
