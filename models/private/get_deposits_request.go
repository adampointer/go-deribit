package private

type GetDepositsRequest struct {
	Count    int64  `json:"count" mapstructure:"count"`
	Currency string `json:"currency" mapstructure:"currency"`
	Offset   int64  `json:"offset" mapstructure:"offset"`
}
