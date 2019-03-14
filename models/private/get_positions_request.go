package private

type GetPositionsRequest struct {
	Currency string `json:"currency" mapstructure:"currency"`
	Kind     string `json:"kind" mapstructure:"kind"`
}
