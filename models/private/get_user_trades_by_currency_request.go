package private

type GetUserTradesByCurrencyRequest struct {
	Count    int64  `json:"count"`
	Currency string `json:"currency"`
	StartID  string `json:"start_id"`
}
