package deribit

// TradeResponse is the data returned by a trade event notification
type TradeResponse struct {
	TradeID       int     `json:"tradeId"`
	Timestamp     int64   `json:"timeStamp"`
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

// OrderBookResponse is the data returned by an orderbook change
type OrderBookResponse struct {
	State           string            `json:"state"`
	SettlementPrice float64           `json:"settlementPrice"`
	Instrument      string            `json:"instrument"`
	Timestamp       int64             `json:"tstamp"`
	Last            float64           `json:"last"`
	Low             float64           `json:"low"`
	High            float64           `json:"high"`
	Mark            float64           `json:"mark"`
	Bids            []*OrderBookEntry `json:"bids"`
	Asks            []*OrderBookEntry `json:"asks"`
}

// OrderBookEntry is an entry in the orderbook
type OrderBookEntry struct {
	Quantity int64   `json:"quantity"`
	Price    float64 `json:"price"`
	Cm       int64   `json:"cm"`
}

// OrderResponse is a response to an OrderRequest
// It contains two fields: the created order in order and a list of the resulting trades in trades
type OrderResponse struct {
	Order  *OrderResponseDetail `json:"order"`
	Trades *OrderResponseTrades `json:"trades"`
}

// OrderResponseTrades trades is populated when a trade is executed immediately
type OrderResponseTrades struct {
	Label      string  `json:"label"`
	SelfTrade  bool    `json:"selfTrade"`
	Quantity   int     `json:"quantity"`
	Price      float64 `json:"price"`
	TradeSeq   int     `json:"tradeSeq"`
	MatchingID string  `json:matchingId`
}

// OrderResponseDetail is the full details of the order
type OrderResponseDetail struct {
	OrderID        string  `json:"orderId"`
	Direction      string  `json:"direction"`
	FilledQuantity int     `json:"filledQuantity"`
	AvgPrice       float64 `json:"avgPrice"`
	Commission     float64 `json:"commission"`
	Created        int64   `json:"created"`
	LastUpdate     int64   `json:"lastUpdate"`
	State          string  `json:"state"`
	API            bool    `json:"api"`
	Triggered      bool    `json:"triggered"`
}
