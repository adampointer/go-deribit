package private

type SubscribeRequest struct {
	Channels []string `json:"channels" mapstructure:"channels"`
}
