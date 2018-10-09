package deribit

// SubscriptionEvent is an event which can be subscribed to
type SubscriptionEvent string

// Events which can be subscribed to
const (
	EvtOrderBook    SubscriptionEvent = "order_book"
	EvtTrade        SubscriptionEvent = "trade"
	EvtMyTrade      SubscriptionEvent = "my_trade"
	EvtUserOrder    SubscriptionEvent = "user_order"
	EvtIndex        SubscriptionEvent = "index"
	EvtPortfolio    SubscriptionEvent = "portfolio"
	EvtAnnouncement SubscriptionEvent = "announcement"
	EvtDelivery     SubscriptionEvent = "delivery"
)

// RPCNotification is a notification which we have subscribed to
type RPCNotification struct {
	Success bool          `json:"success"`
	Testnet bool          `json:"testnet"`
	Message string        `json:"message"`
	Result  []interface{} `json:"result"`
}

// RPCSubscription is a subscription to an event type to receive notifications about
type RPCSubscription struct {
	Data chan []interface{}
}

// EvtTradeResponse is the data returned by a trade event notification
type EvtTradeResponse struct {
	TradeID       int     `json:"tradeId"`
	Timestamp     uint64  `json:"timeStamp"`
	Instrument    string  `json:"instrument"`
	Quantity      int     `json:"quantity"`
	Price         float64 `json:"price"`
	State         string  `json:"state"`
	Direction     string  `json:"direction"`
	OrderID       int     `json:"orderId"`
	MatchingID    int     `json:"matchingId"`
	MakerComm     float64 `json:"makerComm"`
	TakerComm     float64 `json:"takerComm"`
	IndexPrice    float64 `json:"indexPrice"`
	Label         string  `json:"label"`
	Me            string  `json:"me"`
	TickDirection int     `json:"tickDirection"`
}
