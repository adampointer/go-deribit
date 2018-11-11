package deribit

import "strconv"

// OrderRequest is a new order
type OrderRequest struct {
	Instrument  string  `json:"instrument"`
	Quantity    int     `json:"quantity"`
	Type        string  `json:"type"`
	Label       string  `json:"label"`
	Price       float64 `json:"price,omitempty"`
	TimeInForce string  `json:"time_in_force"`
	MaxShow     int     `json:"max_show,omitempty"`
	PostOnly    bool    `json:"post_only,string,omitempty"`
	StopPx      string  `json:"stopPx,omitempty"`
	ExecInst    string  `json:"execInst"`
}

// DefaultMarketOrderRequest returns a limit type OrderRequest with defaults
func DefaultMarketOrderRequest(quantity int) *OrderRequest {
	return &OrderRequest{
		Instrument:  "BTC-PERPETUAL",
		Quantity:    quantity,
		MaxShow:     quantity,
		Type:        "market",
		TimeInForce: "good_til_cancelled",
		PostOnly:    false,
	}
}

// DefaultLimitOrderRequest returns a limit type OrderRequest with defaults
func DefaultLimitOrderRequest(price float64, quantity int, label string) *OrderRequest {
	return &OrderRequest{
		Instrument:  "BTC-PERPETUAL",
		Quantity:    quantity,
		MaxShow:     quantity,
		Type:        "limit",
		Label:       label,
		Price:       price,
		TimeInForce: "good_til_cancelled",
		PostOnly:    true,
	}
}

// DefaultStopOrderRequest returns a stop market type OrderRequest with defaults
func DefaultStopOrderRequest(stopPrice float64, quantity int, label string) *OrderRequest {
	return &OrderRequest{
		Instrument:  "BTC-PERPETUAL",
		Quantity:    quantity,
		MaxShow:     quantity,
		Type:        "stop_market",
		Label:       label,
		StopPx:      strconv.FormatFloat(stopPrice, 'f', -1, 64),
		TimeInForce: "good_til_cancelled",
		ExecInst:    "index_price",
		PostOnly:    false,
	}
}

func (r *OrderRequest) toMap() map[string]interface{} {
	ret := map[string]interface{}{
		"instrument":    r.Instrument,
		"quantity":      r.Quantity,
		"type":          r.Type,
		"label":         r.Label,
		"time_in_force": r.TimeInForce,
		"max_show":      r.MaxShow,
		"post_only":     r.PostOnly,
		"execInst":      r.ExecInst,
	}
	if r.Price > 0 {
		ret["price"] = r.Price
	}
	if len(r.StopPx) > 0 {
		ret["stopPx"] = r.StopPx
	}
	return ret
}
