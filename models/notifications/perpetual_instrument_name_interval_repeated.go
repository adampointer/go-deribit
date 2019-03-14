package notifications

type PerpetualInstrumentNameIntervalRepeated struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  struct {
		Channel string `json:"channel"`
		Data    struct {
			Interest float64 `json:"interest"`
		} `json:"data"`
	} `json:"params"`
}
