package private

type CancelWithdrawalRequest struct {
	Currency string `json:"currency"`
	ID       int64  `json:"id"`
}
