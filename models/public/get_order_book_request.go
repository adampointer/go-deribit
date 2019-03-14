package public

type GetOrderBookRequest struct {
	Depth          int64  `json:"depth"`
	InstrumentName string `json:"instrument_name"`
}
