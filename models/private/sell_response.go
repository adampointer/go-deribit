package private

type SellResponse struct {
	Order struct {
		Amount              int64   `json:"amount"`
		API                 bool    `json:"api"`
		CreationTimestamp   int64   `json:"creation_timestamp"`
		Direction           string  `json:"direction"`
		InstrumentName      string  `json:"instrument_name"`
		IsLiquidation       bool    `json:"is_liquidation"`
		Label               string  `json:"label"`
		LastUpdateTimestamp int64   `json:"last_update_timestamp"`
		MaxShow             int64   `json:"max_show"`
		OrderID             string  `json:"order_id"`
		OrderState          string  `json:"order_state"`
		OrderType           string  `json:"order_type"`
		PostOnly            bool    `json:"post_only"`
		Price               float64 `json:"price"`
		ProfitLoss          int64   `json:"profit_loss"`
		ReduceOnly          bool    `json:"reduce_only"`
		StopPrice           int64   `json:"stop_price"`
		TimeInForce         string  `json:"time_in_force"`
		Trigger             string  `json:"trigger"`
		Triggered           bool    `json:"triggered"`
	} `json:"order"`
	Trades []interface{} `json:"trades"`
}
