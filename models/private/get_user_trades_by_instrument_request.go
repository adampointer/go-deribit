package private

type GetUserTradesByInstrumentRequest struct {
	Count          int64  `json:"count"`
	InstrumentName string `json:"instrument_name"`
	StartSeq       int64  `json:"start_seq"`
}
