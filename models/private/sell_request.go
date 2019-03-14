package private

type SellRequest struct {
	Amount         int64   `json:"amount"`
	InstrumentName string  `json:"instrument_name"`
	Price          float64 `json:"price"`
	StopPrice      int64   `json:"stop_price"`
	Trigger        string  `json:"trigger"`
	Type           string  `json:"type"`
}
