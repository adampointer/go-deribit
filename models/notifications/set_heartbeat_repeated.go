package notifications

type SetHeartbeatRepeated struct {
	Jsonrpc string `json:"jsonrpc" mapstructure:"jsonrpc"`
	Method  string `json:"method" mapstructure:"method"`
	Params  struct {
		Type string `json:"type" mapstructure:"type"`
	} `json:"params" mapstructure:"params"`
}
