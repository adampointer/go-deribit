package private

type GetUserTradesByCurrencyRequest struct {
	Count    int64  `json:"count" mapstructure:"count"`
	Currency string `json:"currency" mapstructure:"currency"`
	StartID  string `json:"start_id" mapstructure:"start_id"`
}
