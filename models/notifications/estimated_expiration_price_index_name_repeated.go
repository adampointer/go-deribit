package notifications

type EstimatedExpirationPriceIndexNameRepeated struct {
	Jsonrpc string `json:"jsonrpc" mapstructure:"jsonrpc"`
	Method  string `json:"method" mapstructure:"method"`
	Params  struct {
		Channel string `json:"channel" mapstructure:"channel"`
		Data    struct {
			IsEstimated bool    `json:"is_estimated" mapstructure:"is_estimated"`
			Price       float64 `json:"price" mapstructure:"price"`
			Seconds     int64   `json:"seconds" mapstructure:"seconds"`
		} `json:"data" mapstructure:"data"`
	} `json:"params" mapstructure:"params"`
}
