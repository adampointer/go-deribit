package notifications

type DeribitPriceIndexIndexNameRepeated struct {
	Jsonrpc string `json:"jsonrpc" mapstructure:"jsonrpc"`
	Method  string `json:"method" mapstructure:"method"`
	Params  struct {
		Channel string `json:"channel" mapstructure:"channel"`
		Data    struct {
			IndexName string  `json:"index_name" mapstructure:"index_name"`
			Price     float64 `json:"price" mapstructure:"price"`
			Timestamp int64   `json:"timestamp" mapstructure:"timestamp"`
		} `json:"data" mapstructure:"data"`
	} `json:"params" mapstructure:"params"`
}
