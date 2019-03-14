package public

type GetLastTradesByInstrumentRequest struct {
	Count          int64  `json:"count"`
	InstrumentName string `json:"instrument_name"`
}
