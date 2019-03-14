package private

type GetUserTradesByInstrumentAndTimeRequest struct {
	EndTimestamp   int64  `json:"end_timestamp"`
	InstrumentName string `json:"instrument_name"`
	StartTimestamp int64  `json:"start_timestamp"`
}
