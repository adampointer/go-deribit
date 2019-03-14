package public

type GetLastTradesByCurrencyRequest struct {
	Count    int64  `json:"count" mapstructure:"count"`
	Currency string `json:"currency" mapstructure:"currency"`
}
