package private

type ClosePositionResponse struct {
	Order struct {
		Amount              int64   `json:"amount" mapstructure:"amount"`
		API                 bool    `json:"api" mapstructure:"api"`
		AveragePrice        float64 `json:"average_price" mapstructure:"average_price"`
		Commission          int64   `json:"commission" mapstructure:"commission"`
		CreationTimestamp   int64   `json:"creation_timestamp" mapstructure:"creation_timestamp"`
		Direction           string  `json:"direction" mapstructure:"direction"`
		FilledAmount        int64   `json:"filled_amount" mapstructure:"filled_amount"`
		InstrumentName      string  `json:"instrument_name" mapstructure:"instrument_name"`
		IsLiquidation       bool    `json:"is_liquidation" mapstructure:"is_liquidation"`
		Label               string  `json:"label" mapstructure:"label"`
		LastUpdateTimestamp int64   `json:"last_update_timestamp" mapstructure:"last_update_timestamp"`
		MaxShow             int64   `json:"max_show" mapstructure:"max_show"`
		OrderID             string  `json:"order_id" mapstructure:"order_id"`
		OrderState          string  `json:"order_state" mapstructure:"order_state"`
		OrderType           string  `json:"order_type" mapstructure:"order_type"`
		PostOnly            bool    `json:"post_only" mapstructure:"post_only"`
		Price               float64 `json:"price" mapstructure:"price"`
		ProfitLoss          int64   `json:"profit_loss" mapstructure:"profit_loss"`
		ReduceOnly          bool    `json:"reduce_only" mapstructure:"reduce_only"`
		TimeInForce         string  `json:"time_in_force" mapstructure:"time_in_force"`
	} `json:"order" mapstructure:"order"`
	Trades []struct {
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
	} `json:"trades" mapstructure:"trades"`
}
