package public

type GetOrderBookRequest struct {
	Depth          int64  `json:"depth" mapstructure:"depth"`
	InstrumentName string `json:"instrument_name" mapstructure:"instrument_name"`
}
