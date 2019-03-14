package private

type BuyRequest struct {
	Amount         int64  `json:"amount"`
	InstrumentName string `json:"instrument_name"`
	Label          string `json:"label"`
	Type           string `json:"type"`
}
