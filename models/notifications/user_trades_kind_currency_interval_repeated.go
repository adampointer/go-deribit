package notifications

type UserTradesKindCurrencyIntervalRepeated struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  struct {
		Channel string `json:"channel"`
		Data    []struct {
			Amount         int64       `json:"amount"`
			Direction      string      `json:"direction"`
			Fee            int64       `json:"fee"`
			FeeCurrency    string      `json:"fee_currency"`
			IndexPrice     float64     `json:"index_price"`
			InstrumentName string      `json:"instrument_name"`
			Liquidity      string      `json:"liquidity"`
			MatchingID     interface{} `json:"matching_id"`
			OrderID        string      `json:"order_id"`
			OrderType      string      `json:"order_type"`
			Price          float64     `json:"price"`
			SelfTrade      bool        `json:"self_trade"`
			State          string      `json:"state"`
			TickDirection  int64       `json:"tick_direction"`
			Timestamp      int64       `json:"timestamp"`
			TradeID        string      `json:"trade_id"`
			TradeSeq       int64       `json:"trade_seq"`
		} `json:"data"`
	} `json:"params"`
}
