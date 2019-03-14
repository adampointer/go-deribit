package public

type TickerResponse struct {
	BestAskAmount   int64   `json:"best_ask_amount" mapstructure:"best_ask_amount"`
	BestAskPrice    float64 `json:"best_ask_price" mapstructure:"best_ask_price"`
	BestBidAmount   int64   `json:"best_bid_amount" mapstructure:"best_bid_amount"`
	BestBidPrice    float64 `json:"best_bid_price" mapstructure:"best_bid_price"`
	CurrentFunding  int64   `json:"current_funding" mapstructure:"current_funding"`
	Funding8h       float64 `json:"funding_8h" mapstructure:"funding_8h"`
	IndexPrice      float64 `json:"index_price" mapstructure:"index_price"`
	InstrumentName  string  `json:"instrument_name" mapstructure:"instrument_name"`
	LastPrice       float64 `json:"last_price" mapstructure:"last_price"`
	MarkPrice       float64 `json:"mark_price" mapstructure:"mark_price"`
	MaxPrice        float64 `json:"max_price" mapstructure:"max_price"`
	MinPrice        float64 `json:"min_price" mapstructure:"min_price"`
	OpenInterest    float64 `json:"open_interest" mapstructure:"open_interest"`
	SettlementPrice float64 `json:"settlement_price" mapstructure:"settlement_price"`
	State           string  `json:"state" mapstructure:"state"`
	Stats           struct {
		High   float64 `json:"high" mapstructure:"high"`
		Low    int64   `json:"low" mapstructure:"low"`
		Volume float64 `json:"volume" mapstructure:"volume"`
	} `json:"stats" mapstructure:"stats"`
	Timestamp int64 `json:"timestamp" mapstructure:"timestamp"`
}
