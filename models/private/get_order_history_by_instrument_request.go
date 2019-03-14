package private

type GetOrderHistoryByInstrumentRequest struct {
	Count          int64  `json:"count"`
	InstrumentName string `json:"instrument_name"`
}
