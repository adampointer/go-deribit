package private

type GetUserTradesByInstrumentRequest struct {
	Count          int64  `json:"count" mapstructure:"count"`
	InstrumentName string `json:"instrument_name" mapstructure:"instrument_name"`
	StartSeq       int64  `json:"start_seq" mapstructure:"start_seq"`
}
