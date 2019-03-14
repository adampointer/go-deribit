package private

type GetPositionRequest struct {
	InstrumentName string `json:"instrument_name" mapstructure:"instrument_name"`
}
