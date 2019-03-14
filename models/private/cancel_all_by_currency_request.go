package private

type CancelAllByCurrencyRequest struct {
	Currency string `json:"currency"`
	Kind     string `json:"kind"`
}
