package notifications

type UserTradesInstrumentNameIntervalRepeated struct {
	Jsonrpc string `json:"jsonrpc" mapstructure:"jsonrpc"`
	Method  string `json:"method" mapstructure:"method"`
	Params  struct {
		Channel string `json:"channel" mapstructure:"channel"`
		Data    []struct {
			Amount         int64       `json:"amount" mapstructure:"amount"`
			Direction      string      `json:"direction" mapstructure:"direction"`
			Fee            int64       `json:"fee" mapstructure:"fee"`
			FeeCurrency    string      `json:"fee_currency" mapstructure:"fee_currency"`
			IndexPrice     float64     `json:"index_price" mapstructure:"index_price"`
			InstrumentName string      `json:"instrument_name" mapstructure:"instrument_name"`
			Liquidity      string      `json:"liquidity" mapstructure:"liquidity"`
			MatchingID     interface{} `json:"matching_id" mapstructure:"matching_id"`
			OrderID        string      `json:"order_id" mapstructure:"order_id"`
			OrderType      string      `json:"order_type" mapstructure:"order_type"`
			Price          float64     `json:"price" mapstructure:"price"`
			SelfTrade      bool        `json:"self_trade" mapstructure:"self_trade"`
			State          string      `json:"state" mapstructure:"state"`
			TickDirection  int64       `json:"tick_direction" mapstructure:"tick_direction"`
			Timestamp      int64       `json:"timestamp" mapstructure:"timestamp"`
			TradeID        string      `json:"trade_id" mapstructure:"trade_id"`
			TradeSeq       int64       `json:"trade_seq" mapstructure:"trade_seq"`
		} `json:"data" mapstructure:"data"`
	} `json:"params" mapstructure:"params"`
}
