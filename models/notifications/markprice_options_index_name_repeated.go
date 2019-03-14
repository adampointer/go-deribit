package notifications

type MarkpriceOptionsIndexNameRepeated struct {
	Jsonrpc string `json:"jsonrpc" mapstructure:"jsonrpc"`
	Method  string `json:"method" mapstructure:"method"`
	Params  struct {
		Channel string `json:"channel" mapstructure:"channel"`
		Data    []struct {
			InstrumentName string  `json:"instrument_name" mapstructure:"instrument_name"`
			Iv             float64 `json:"iv" mapstructure:"iv"`
			MarkPrice      float64 `json:"mark_price" mapstructure:"mark_price"`
		} `json:"data" mapstructure:"data"`
	} `json:"params" mapstructure:"params"`
}
