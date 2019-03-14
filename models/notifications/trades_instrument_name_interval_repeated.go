package notifications

type TradesInstrumentNameIntervalRepeated struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  struct {
		Channel string `json:"channel"`
		Data    []struct {
			Amount         int64   `json:"amount"`
			Direction      string  `json:"direction"`
			IndexPrice     float64 `json:"index_price"`
			InstrumentName string  `json:"instrument_name"`
			Price          float64 `json:"price"`
			TickDirection  int64   `json:"tick_direction"`
			Timestamp      int64   `json:"timestamp"`
			TradeID        string  `json:"trade_id"`
			TradeSeq       int64   `json:"trade_seq"`
		} `json:"data"`
	} `json:"params"`
}
