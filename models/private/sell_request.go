package private

type SellRequest struct {
	Amount         int64   `json:"amount" mapstructure:"amount"`
	InstrumentName string  `json:"instrument_name" mapstructure:"instrument_name"`
	Price          float64 `json:"price" mapstructure:"price"`
	StopPrice      int64   `json:"stop_price" mapstructure:"stop_price"`
	Trigger        string  `json:"trigger" mapstructure:"trigger"`
	Type           string  `json:"type" mapstructure:"type"`
}
