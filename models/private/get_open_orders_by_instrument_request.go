package private

type GetOpenOrdersByInstrumentRequest struct {
	InstrumentName string `json:"instrument_name" mapstructure:"instrument_name"`
}
