package private

type GetSettlementHistoryByInstrumentRequest struct {
	Count          int64  `json:"count" mapstructure:"count"`
	InstrumentName string `json:"instrument_name" mapstructure:"instrument_name"`
	Type           string `json:"type" mapstructure:"type"`
}
