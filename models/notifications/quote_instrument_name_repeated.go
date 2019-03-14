package notifications

type QuoteInstrumentNameRepeated struct {
	Jsonrpc string `json:"jsonrpc" mapstructure:"jsonrpc"`
	Method  string `json:"method" mapstructure:"method"`
	Params  struct {
		Channel string `json:"channel" mapstructure:"channel"`
		Data    struct {
			BestAskAmount  int64   `json:"best_ask_amount" mapstructure:"best_ask_amount"`
			BestAskPrice   float64 `json:"best_ask_price" mapstructure:"best_ask_price"`
			BestBidAmount  int64   `json:"best_bid_amount" mapstructure:"best_bid_amount"`
			BestBidPrice   float64 `json:"best_bid_price" mapstructure:"best_bid_price"`
			InstrumentName string  `json:"instrument_name" mapstructure:"instrument_name"`
			Timestamp      int64   `json:"timestamp" mapstructure:"timestamp"`
		} `json:"data" mapstructure:"data"`
	} `json:"params" mapstructure:"params"`
}
