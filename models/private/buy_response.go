package private

type BuyResponse struct {
	Order struct {
		Amount              int64   `json:"amount"`
		API                 bool    `json:"api"`
		AveragePrice        float64 `json:"average_price"`
		Commission          float64 `json:"commission"`
		CreationTimestamp   int64   `json:"creation_timestamp"`
		Direction           string  `json:"direction"`
		FilledAmount        int64   `json:"filled_amount"`
		InstrumentName      string  `json:"instrument_name"`
		IsLiquidation       bool    `json:"is_liquidation"`
		Label               string  `json:"label"`
		LastUpdateTimestamp int64   `json:"last_update_timestamp"`
		MaxShow             int64   `json:"max_show"`
		OrderID             string  `json:"order_id"`
		OrderState          string  `json:"order_state"`
		OrderType           string  `json:"order_type"`
		PostOnly            bool    `json:"post_only"`
		Price               string  `json:"price"`
		ProfitLoss          int64   `json:"profit_loss"`
		ReduceOnly          bool    `json:"reduce_only"`
		TimeInForce         string  `json:"time_in_force"`
	} `json:"order"`
	Trades []struct {
		Amount         int64       `json:"amount"`
		Direction      string      `json:"direction"`
		Fee            float64     `json:"fee"`
		FeeCurrency    string      `json:"fee_currency"`
		IndexPrice     float64     `json:"index_price"`
		InstrumentName string      `json:"instrument_name"`
		Label          string      `json:"label"`
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
	} `json:"trades"`
}
