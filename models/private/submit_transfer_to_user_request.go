package private

type SubmitTransferToUserRequest struct {
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`
	Destination string  `json:"destination"`
}
