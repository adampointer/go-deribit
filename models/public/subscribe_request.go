package public

type SubscribeRequest struct {
	Channels []string `json:"channels" mapstructure:"channels"`
}
