package notifications

type BookInstrumentNameGroupDepthIntervalRepeated struct {
	Jsonrpc string `json:"jsonrpc" mapstructure:"jsonrpc"`
	Method  string `json:"method" mapstructure:"method"`
	Params  struct {
		Channel string `json:"channel" mapstructure:"channel"`
		Data    struct {
			Asks           [][]int64 `json:"asks" mapstructure:"asks"`
			Bids           [][]int64 `json:"bids" mapstructure:"bids"`
			ChangeID       int64     `json:"change_id" mapstructure:"change_id"`
			InstrumentName string    `json:"instrument_name" mapstructure:"instrument_name"`
			Timestamp      int64     `json:"timestamp" mapstructure:"timestamp"`
		} `json:"data" mapstructure:"data"`
	} `json:"params" mapstructure:"params"`
}
