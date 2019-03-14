package private

type GetOrderMarginByIdsResponse []struct {
	InitialMargin float64 `json:"initial_margin"`
	OrderID       string  `json:"order_id"`
}
