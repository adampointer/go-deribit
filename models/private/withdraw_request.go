package private

type WithdrawRequest struct {
	Address  string  `json:"address" mapstructure:"address"`
	Amount   float64 `json:"amount" mapstructure:"amount"`
	Currency string  `json:"currency" mapstructure:"currency"`
	Priority string  `json:"priority" mapstructure:"priority"`
}
