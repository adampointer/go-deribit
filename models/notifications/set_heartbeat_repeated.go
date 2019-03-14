package notifications

type SetHeartbeatRepeated struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  struct {
		Type string `json:"type"`
	} `json:"params"`
}
