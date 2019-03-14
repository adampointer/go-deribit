package private

type GetAccountSummaryRequest struct {
	Currency string `json:"currency"`
	Extended bool   `json:"extended"`
}
