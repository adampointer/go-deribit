package notifications

type DeribitPriceRankingIndexNameRepeated struct {
	Jsonrpc string `json:"jsonrpc" mapstructure:"jsonrpc"`
	Method  string `json:"method" mapstructure:"method"`
	Params  struct {
		Channel string `json:"channel" mapstructure:"channel"`
		Data    []struct {
			Enabled    bool    `json:"enabled" mapstructure:"enabled"`
			Identifier string  `json:"identifier" mapstructure:"identifier"`
			Price      float64 `json:"price" mapstructure:"price"`
			Timestamp  int64   `json:"timestamp" mapstructure:"timestamp"`
			Weight     int64   `json:"weight" mapstructure:"weight"`
		} `json:"data" mapstructure:"data"`
	} `json:"params" mapstructure:"params"`
}
