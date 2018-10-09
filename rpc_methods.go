package deribit

import "fmt"

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

// Subscribe to notifications about an event
func (e *Exchange) Subscribe(event SubscriptionEvent, c chan []interface{}) error {
	req := &RPCRequest{
		Action: "/api/v1/private/subscribe",
		Arguments: map[string]interface{}{
			"event":      []SubscriptionEvent{event},
			"continue":   true,
			"instrument": []string{"BTC-PERPETUAL"},
		},
	}
	req.GenerateSig(e.key, e.secret)
	res, err := e.makeRequest(*req)
	if err != nil {
		return err
	}
	if !res.Success {
		return fmt.Errorf("Subscription to %s was not successful: %s", event, res.Message)
	}
	sub := &RPCSubscription{Data: c}
	e.subscriptions[string(event)+"_event"] = sub
	return nil
}
