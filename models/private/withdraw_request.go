package private

type WithdrawRequest struct {
	Address  string  `json:"address"`
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
	Priority string  `json:"priority"`
}
