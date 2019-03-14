package notifications

type BookInstrumentNameIntervalRepeated struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  struct {
		Channel string `json:"channel"`
		Data    struct {
			Asks             [][]interface{} `json:"asks"`
			Bids             []interface{}   `json:"bids"`
			ChangeID         int64           `json:"change_id"`
			InstrumentName   string          `json:"instrument_name"`
			PreviousChangeID int64           `json:"previous_change_id"`
			Timestamp        int64           `json:"timestamp"`
		} `json:"data"`
	} `json:"params"`
}
