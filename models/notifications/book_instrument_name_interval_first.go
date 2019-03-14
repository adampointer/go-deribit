package notifications

type BookInstrumentNameIntervalFirst struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  struct {
		Channel string `json:"channel"`
		Data    struct {
			Asks           []interface{} `json:"asks"`
			Bids           [][]int64     `json:"bids"`
			ChangeID       int64         `json:"change_id"`
			InstrumentName string        `json:"instrument_name"`
			Timestamp      int64         `json:"timestamp"`
		} `json:"data"`
	} `json:"params"`
}
