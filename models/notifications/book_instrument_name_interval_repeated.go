package notifications

type BookInstrumentNameIntervalRepeated struct {
	Jsonrpc string `json:"jsonrpc" mapstructure:"jsonrpc"`
	Method  string `json:"method" mapstructure:"method"`
	Params  struct {
		Channel string `json:"channel" mapstructure:"channel"`
		Data    struct {
			Asks             [][]interface{} `json:"asks" mapstructure:"asks"`
			Bids             []interface{}   `json:"bids" mapstructure:"bids"`
			ChangeID         int64           `json:"change_id" mapstructure:"change_id"`
			InstrumentName   string          `json:"instrument_name" mapstructure:"instrument_name"`
			PreviousChangeID int64           `json:"previous_change_id" mapstructure:"previous_change_id"`
			Timestamp        int64           `json:"timestamp" mapstructure:"timestamp"`
		} `json:"data" mapstructure:"data"`
	} `json:"params" mapstructure:"params"`
}
