package private

type GetUserTradesByOrderRequest struct {
	OrderID string `json:"order_id" mapstructure:"order_id"`
}
