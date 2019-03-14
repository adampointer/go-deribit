package private

type GetSettlementHistoryByInstrumentRequest struct {
	Count          int64  `json:"count"`
	InstrumentName string `json:"instrument_name"`
	Type           string `json:"type"`
}
