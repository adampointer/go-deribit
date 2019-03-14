package notifications

type DeribitPriceRankingIndexNameRepeated struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  struct {
		Channel string `json:"channel"`
		Data    []struct {
			Enabled    bool    `json:"enabled"`
			Identifier string  `json:"identifier"`
			Price      float64 `json:"price"`
			Timestamp  int64   `json:"timestamp"`
			Weight     int64   `json:"weight"`
		} `json:"data"`
	} `json:"params"`
}
