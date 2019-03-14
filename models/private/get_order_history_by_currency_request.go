package private

type GetOrderHistoryByCurrencyRequest struct {
	Count    int64  `json:"count"`
	Currency string `json:"currency"`
	Kind     string `json:"kind"`
}
