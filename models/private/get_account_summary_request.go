package private

type GetAccountSummaryRequest struct {
	Currency string `json:"currency" mapstructure:"currency"`
	Extended bool   `json:"extended" mapstructure:"extended"`
}
