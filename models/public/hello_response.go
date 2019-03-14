package public

type HelloResponse struct {
	Version string `json:"version" mapstructure:"version"`
}
