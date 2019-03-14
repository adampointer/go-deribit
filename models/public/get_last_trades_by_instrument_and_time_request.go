package public

type GetLastTradesByInstrumentAndTimeRequest struct {
	Count          int64  `json:"count"`
	EndTimestamp   int64  `json:"end_timestamp"`
	InstrumentName string `json:"instrument_name"`
}
