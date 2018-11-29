package deribit

import (
	"encoding/json"
	"fmt"
)

// SubscriptionEvent is an event which can be subscribed to
type SubscriptionEvent string

// Events which can be subscribed to
const (
	evtOrderBook    SubscriptionEvent = "order_book"
	evtTrade        SubscriptionEvent = "trade"
	evtMyTrade      SubscriptionEvent = "my_trade"
	evtUserOrder    SubscriptionEvent = "user_order"
	evtIndex        SubscriptionEvent = "index"
	evtPortfolio    SubscriptionEvent = "portfolio"
	evtAnnouncement SubscriptionEvent = "announcement"
	evtDelivery     SubscriptionEvent = "delivery"
)

// SubscribeTrades subscribes to trade notifications which are sent on the returned channel
func (e *Exchange) SubscribeTrades(instrument string) (chan *RPCNotification, error) {
	c := make(chan *RPCNotification)
	if err := e.subscribe(evtTrade, c, instrument); err != nil {
		return c, err
	}
	return c, nil
}

// SubscribeOrderBook subscribes to orderbook notifications which are sent on the returned channel
func (e *Exchange) SubscribeOrderBook(instrument string) (chan *RPCNotification, error) {
	c := make(chan *RPCNotification)
	if err := e.subscribe(evtOrderBook, c, instrument); err != nil {
		return c, err
	}
	return c, nil
}

// SubscribeUserOrders subscribes to changes of your orders
func (e *Exchange) SubscribeUserOrders(instrument string) (chan *RPCNotification, error) {
	c := make(chan *RPCNotification)
	if err := e.subscribe(evtUserOrder, c, instrument); err != nil {
		return c, err
	}
	return c, nil
}

// subscribe to notifications about an event
func (e *Exchange) subscribe(event SubscriptionEvent, c chan *RPCNotification, instrument string) error {
	req := &RPCRequest{
		Action: "/api/v1/private/subscribe",
		Arguments: map[string]interface{}{
			"event":      []SubscriptionEvent{event},
			"continue":   true,
			"instrument": []string{instrument},
		},
	}
	if err := req.GenerateSig(e.key, e.secret); err != nil {
		return err
	}
	res, err := e.makeRequest(*req)
	if err != nil {
		return err
	}
	if !res.Success {
		return fmt.Errorf("Subscription to %s was not successful: %s", event, res.Message)
	}
	sub := &RPCSubscription{Data: c, Type: string(event)}
	// Ok well this is annoying. If I find other exceptions where the message field
	// does not exactly match the event name I will have to make something more robust
	// than this little hack
	evtType := string(event)
	if evtType == "user_order" {
		evtType = "user_orders"
	}
	e.subscriptions[evtType+"_event"] = sub
	return nil
}

// RPCSubscription is a subscription to an event type to receive notifications about
type RPCSubscription struct {
	Data chan *RPCNotification
	Type string
}

// RPCNotification is a notification which we have subscribed to
type RPCNotification struct {
	Success bool            `json:"success"`
	Testnet bool            `json:"testnet"`
	Message string          `json:"message"`
	Result  json.RawMessage `json:"result"`
}

// DecodeTrades takes the result from the notification response and decodes it into a slice of trade responses
func (n *RPCNotification) DecodeTrades() ([]*TradeResponse, error) {
	if n.Message != fmt.Sprintf("%s_event", evtTrade) {
		return nil, fmt.Errorf("Attempt to convert %s notification to trades", n.Message)
	}
	var ret []*TradeResponse
	if err := json.Unmarshal(n.Result, &ret); err != nil {
		return nil, fmt.Errorf("Unable to unmarshal result: %s", err)
	}
	return ret, nil
}

// DecodeOrderBook takes the result from the notification response and decodes it an orderbook response
func (n *RPCNotification) DecodeOrderBook() (*OrderBookResponse, error) {
	if n.Message != fmt.Sprintf("%s_event", evtOrderBook) {
		return nil, fmt.Errorf("Attempt to convert %s notification to an order book", n.Message)
	}
	var ret OrderBookResponse
	if err := json.Unmarshal(n.Result, &ret); err != nil {
		return nil, fmt.Errorf("Unable to unmarshal result: %s", err)
	}
	return &ret, nil
}

// DecodeUserOrder takes the result from the notification response and decodes it an order response
func (n *RPCNotification) DecodeUserOrder() ([]*OrderResponseDetail, error) {
	if n.Message != fmt.Sprintf("user_orders_event") {
		return nil, fmt.Errorf("Attempt to convert %s notification to a user order", n.Message)
	}
	var ret []*OrderResponseDetail
	if err := json.Unmarshal(n.Result, &ret); err != nil {
		return nil, fmt.Errorf("Unable to unmarshal result: %s", err)
	}
	return ret, nil
}
