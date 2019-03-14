package public

type GetLastTradesByCurrencyResponse struct {
	HasMore bool `json:"has_more"`
	Trades  []struct {
		Amount         int64   `json:"amount"`
		Direction      string  `json:"direction"`
		IndexPrice     float64 `json:"index_price"`
		InstrumentName string  `json:"instrument_name"`
		Iv             float64 `json:"iv"`
		Price          float64 `json:"price"`
		TickDirection  int64   `json:"tick_direction"`
		Timestamp      int64   `json:"timestamp"`
		TradeID        string  `json:"trade_id"`
		TradeSeq       int64   `json:"trade_seq"`
	} `json:"trades"`
}
