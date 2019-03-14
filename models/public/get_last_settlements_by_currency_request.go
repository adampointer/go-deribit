package public

type GetLastSettlementsByCurrencyRequest struct {
	Count    int64  `json:"count"`
	Currency string `json:"currency"`
	Type     string `json:"type"`
}
