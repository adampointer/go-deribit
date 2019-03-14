package public

type GetInstrumentsRequest struct {
	Currency string `json:"currency"`
	Expired  bool   `json:"expired"`
	Kind     string `json:"kind"`
}
