package private

type SubmitTransferToSubaccountRequest struct {
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`
	Destination int64   `json:"destination"`
}
