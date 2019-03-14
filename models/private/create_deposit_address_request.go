package private

type CreateDepositAddressRequest struct {
	Currency string `json:"currency" mapstructure:"currency"`
}
