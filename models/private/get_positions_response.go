package private

type GetPositionsResponse []struct {
	AveragePrice              int64   `json:"average_price"`
	Delta                     int64   `json:"delta"`
	Direction                 string  `json:"direction"`
	EstimatedLiquidationPrice int64   `json:"estimated_liquidation_price"`
	FloatingProfitLoss        int64   `json:"floating_profit_loss"`
	IndexPrice                float64 `json:"index_price"`
	InitialMargin             int64   `json:"initial_margin"`
	InstrumentName            string  `json:"instrument_name"`
	Kind                      string  `json:"kind"`
	MaintenanceMargin         int64   `json:"maintenance_margin"`
	MarkPrice                 float64 `json:"mark_price"`
	OpenOrdersMargin          float64 `json:"open_orders_margin"`
	RealizedProfitLoss        int64   `json:"realized_profit_loss"`
	SettlementPrice           float64 `json:"settlement_price"`
	Size                      int64   `json:"size"`
	SizeCurrency              int64   `json:"size_currency"`
	TotalProfitLoss           int64   `json:"total_profit_loss"`
}
