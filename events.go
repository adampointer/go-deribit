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

// RPCNotification is a notification which we have subscribed to
type RPCNotification struct {
	Success bool            `json:"success"`
	Testnet bool            `json:"testnet"`
	Message string          `json:"message"`
	Result  json.RawMessage `json:"result"`
}

// RPCSubscription is a subscription to an event type to receive notifications about
type RPCSubscription struct {
	Data chan []interface{}
	Type string
}

// SubscribeTrades subscribes to trade notifications which are sent on the returned channel
// You will need to cast interface{} to *TradeResponse the other end
func (e *Exchange) SubscribeTrades() (chan []interface{}, error) {
	c := make(chan []interface{})
	if err := e.subscribe(evtTrade, c); err != nil {
		return c, err
	}
	return c, nil
}

// subscribe to notifications about an event
func (e *Exchange) subscribe(event SubscriptionEvent, c chan []interface{}) error {
	req := &RPCRequest{
		Action: "/api/v1/private/subscribe",
		Arguments: map[string]interface{}{
			"event":      []SubscriptionEvent{event},
			"continue":   true,
			"instrument": []string{"BTC-PERPETUAL"},
		},
	}
	req.GenerateSig(e.key, e.secret)
	res, err := e.makeRequest(*req, "")
	if err != nil {
		return err
	}
	if !res.Success {
		return fmt.Errorf("Subscription to %s was not successful: %s", event, res.Message)
	}
	sub := &RPCSubscription{Data: c, Type: string(event)}
	e.subscriptions[string(event)+"_event"] = sub
	return nil
}
