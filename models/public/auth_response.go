package public

type AuthResponse struct {
	AccessToken  string `json:"access_token" mapstructure:"access_token"`
	ExpiresIn    int64  `json:"expires_in" mapstructure:"expires_in"`
	RefreshToken string `json:"refresh_token" mapstructure:"refresh_token"`
	Scope        string `json:"scope" mapstructure:"scope"`
	TokenType    string `json:"token_type" mapstructure:"token_type"`
}
