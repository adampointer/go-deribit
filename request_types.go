package deribit

import "encoding/json"

// OrderRequest is a new order
type OrderRequest struct {
	Instrument  string  `json:"instrument"`
	Quantity    int     `json:"quantity"`
	Type        string  `json:"type"`
	Label       string  `json:"label"`
	Price       float64 `json:"price"`
	TimeInForce string  `json:"time_in_force"`
	MaxShow     int     `json:"max_show"`
	PostOnly    bool    `json:"post_only,string"`
	StopPx      string  `json:"stopPx,omitempty"`
	ExecInst    string  `json:"execInst"`
}

// DefaultLimitOrderRequest returns a limit type OrderRequest with defaults
func DefaultLimitOrderRequest(price float64, quantity int, label string) *OrderRequest {
	return &OrderRequest{
		Instrument:  "BTC-PERPETUAL",
		Quantity:    quantity,
		Type:        "limit",
		Label:       label,
		Price:       price,
		TimeInForce: "good_til_cancelled",
		MaxShow:     1,
		PostOnly:    true,
	}
}

func (r *OrderRequest) toMap() (map[string]interface{}, error) {
	var out map[string]interface{}
	inrec, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(inrec, out); err != nil {
		return nil, err
	}
	return out, nil
}
