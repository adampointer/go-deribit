package public

type TickerRequest struct {
	InstrumentName string `json:"instrument_name" mapstructure:"instrument_name"`
}
