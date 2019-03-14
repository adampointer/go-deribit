package notifications

type PerpetualInstrumentNameIntervalRepeated struct {
	Jsonrpc string `json:"jsonrpc" mapstructure:"jsonrpc"`
	Method  string `json:"method" mapstructure:"method"`
	Params  struct {
		Channel string `json:"channel" mapstructure:"channel"`
		Data    struct {
			Interest float64 `json:"interest" mapstructure:"interest"`
		} `json:"data" mapstructure:"data"`
	} `json:"params" mapstructure:"params"`
}
