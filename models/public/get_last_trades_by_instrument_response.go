package public

type GetLastTradesByInstrumentResponse struct {
	HasMore bool `json:"has_more" mapstructure:"has_more"`
	Trades  []struct {
		Amount         int64   `json:"amount" mapstructure:"amount"`
		Direction      string  `json:"direction" mapstructure:"direction"`
		IndexPrice     float64 `json:"index_price" mapstructure:"index_price"`
		InstrumentName string  `json:"instrument_name" mapstructure:"instrument_name"`
		Price          int64   `json:"price" mapstructure:"price"`
		TickDirection  int64   `json:"tick_direction" mapstructure:"tick_direction"`
		Timestamp      int64   `json:"timestamp" mapstructure:"timestamp"`
		TradeID        string  `json:"trade_id" mapstructure:"trade_id"`
		TradeSeq       int64   `json:"trade_seq" mapstructure:"trade_seq"`
	} `json:"trades" mapstructure:"trades"`
}
