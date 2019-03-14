package private

type GetOrderHistoryByCurrencyRequest struct {
	Count    int64  `json:"count" mapstructure:"count"`
	Currency string `json:"currency" mapstructure:"currency"`
	Kind     string `json:"kind" mapstructure:"kind"`
}
