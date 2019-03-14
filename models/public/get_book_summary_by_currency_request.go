package public

type GetBookSummaryByCurrencyRequest struct {
	Currency string `json:"currency"`
	Kind     string `json:"kind"`
}
