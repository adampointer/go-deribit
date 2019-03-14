package public

type GetInstrumentsRequest struct {
	Currency string `json:"currency" mapstructure:"currency"`
	Expired  bool   `json:"expired" mapstructure:"expired"`
	Kind     string `json:"kind" mapstructure:"kind"`
}
