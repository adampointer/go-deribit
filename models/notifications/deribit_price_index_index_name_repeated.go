package notifications

type DeribitPriceIndexIndexNameRepeated struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  struct {
		Channel string `json:"channel"`
		Data    struct {
			IndexName string  `json:"index_name"`
			Price     float64 `json:"price"`
			Timestamp int64   `json:"timestamp"`
		} `json:"data"`
	} `json:"params"`
}
