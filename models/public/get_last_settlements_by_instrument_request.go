package public

type GetLastSettlementsByInstrumentRequest struct {
	Count          int64  `json:"count"`
	InstrumentName string `json:"instrument_name"`
	Type           string `json:"type"`
}
