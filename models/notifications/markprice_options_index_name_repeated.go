package notifications

type MarkpriceOptionsIndexNameRepeated struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  struct {
		Channel string `json:"channel"`
		Data    []struct {
			InstrumentName string  `json:"instrument_name"`
			Iv             float64 `json:"iv"`
			MarkPrice      float64 `json:"mark_price"`
		} `json:"data"`
	} `json:"params"`
}
