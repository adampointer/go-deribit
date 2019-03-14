package private

type GetWithdrawalsRequest struct {
	Count    int64  `json:"count"`
	Currency string `json:"currency"`
	Offset   int64  `json:"offset"`
}
