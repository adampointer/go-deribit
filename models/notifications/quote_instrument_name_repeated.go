package notifications

type QuoteInstrumentNameRepeated struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  struct {
		Channel string `json:"channel"`
		Data    struct {
			BestAskAmount  int64   `json:"best_ask_amount"`
			BestAskPrice   float64 `json:"best_ask_price"`
			BestBidAmount  int64   `json:"best_bid_amount"`
			BestBidPrice   float64 `json:"best_bid_price"`
			InstrumentName string  `json:"instrument_name"`
			Timestamp      int64   `json:"timestamp"`
		} `json:"data"`
	} `json:"params"`
}
