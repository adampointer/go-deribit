package notifications

type TradesInstrumentNameIntervalRepeated struct {
	Jsonrpc string `json:"jsonrpc" mapstructure:"jsonrpc"`
	Method  string `json:"method" mapstructure:"method"`
	Params  struct {
		Channel string `json:"channel" mapstructure:"channel"`
		Data    []struct {
			Amount         int64   `json:"amount" mapstructure:"amount"`
			Direction      string  `json:"direction" mapstructure:"direction"`
			IndexPrice     float64 `json:"index_price" mapstructure:"index_price"`
			InstrumentName string  `json:"instrument_name" mapstructure:"instrument_name"`
			Price          float64 `json:"price" mapstructure:"price"`
			TickDirection  int64   `json:"tick_direction" mapstructure:"tick_direction"`
			Timestamp      int64   `json:"timestamp" mapstructure:"timestamp"`
			TradeID        string  `json:"trade_id" mapstructure:"trade_id"`
			TradeSeq       int64   `json:"trade_seq" mapstructure:"trade_seq"`
		} `json:"data" mapstructure:"data"`
	} `json:"params" mapstructure:"params"`
}
