package private

type EditRequest struct {
	Advanced string `json:"advanced"`
	Amount   int64  `json:"amount"`
	OrderID  string `json:"order_id"`
	Price    int64  `json:"price"`
}
