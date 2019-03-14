package public

type UnsubscribeRequest struct {
	Channels []string `json:"channels" mapstructure:"channels"`
}
