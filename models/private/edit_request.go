package private

type EditRequest struct {
	Advanced string `json:"advanced" mapstructure:"advanced"`
	Amount   int64  `json:"amount" mapstructure:"amount"`
	OrderID  string `json:"order_id" mapstructure:"order_id"`
	Price    int64  `json:"price" mapstructure:"price"`
}
