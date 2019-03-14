package public

type GetLastTradesByInstrumentAndTimeRequest struct {
	Count          int64  `json:"count" mapstructure:"count"`
	EndTimestamp   int64  `json:"end_timestamp" mapstructure:"end_timestamp"`
	InstrumentName string `json:"instrument_name" mapstructure:"instrument_name"`
}
