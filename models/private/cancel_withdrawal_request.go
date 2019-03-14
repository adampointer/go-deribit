package private

type CancelWithdrawalRequest struct {
	Currency string `json:"currency" mapstructure:"currency"`
	ID       int64  `json:"id" mapstructure:"id"`
}
