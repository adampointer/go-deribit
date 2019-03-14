package private

type UnsubscribeRequest struct {
	Channels []string `json:"channels" mapstructure:"channels"`
}
