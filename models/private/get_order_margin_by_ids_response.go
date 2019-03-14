package private

type GetOrderMarginByIdsResponse []struct {
	InitialMargin float64 `json:"initial_margin" mapstructure:"initial_margin"`
	OrderID       string  `json:"order_id" mapstructure:"order_id"`
}
