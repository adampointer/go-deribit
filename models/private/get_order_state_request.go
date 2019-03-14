package private

type GetOrderStateRequest struct {
	OrderID string `json:"order_id" mapstructure:"order_id"`
}
