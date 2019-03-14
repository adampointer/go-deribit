package notifications

type EstimatedExpirationPriceIndexNameRepeated struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  struct {
		Channel string `json:"channel"`
		Data    struct {
			IsEstimated bool    `json:"is_estimated"`
			Price       float64 `json:"price"`
			Seconds     int64   `json:"seconds"`
		} `json:"data"`
	} `json:"params"`
}
