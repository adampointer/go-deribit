package public

type GetLastTradesByCurrencyRequest struct {
	Count    int64  `json:"count"`
	Currency string `json:"currency"`
}
