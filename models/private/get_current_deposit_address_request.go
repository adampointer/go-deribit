package private

type GetCurrentDepositAddressRequest struct {
	Currency string `json:"currency" mapstructure:"currency"`
}
