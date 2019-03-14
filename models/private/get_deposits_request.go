package private

type GetDepositsRequest struct {
	Count    int64  `json:"count"`
	Currency string `json:"currency"`
	Offset   int64  `json:"offset"`
}
