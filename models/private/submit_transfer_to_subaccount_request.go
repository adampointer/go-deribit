package private

type SubmitTransferToSubaccountRequest struct {
	Amount      float64 `json:"amount" mapstructure:"amount"`
	Currency    string  `json:"currency" mapstructure:"currency"`
	Destination int64   `json:"destination" mapstructure:"destination"`
}
