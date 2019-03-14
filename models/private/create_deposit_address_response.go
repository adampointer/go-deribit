package private

type CreateDepositAddressResponse struct {
	Address           string `json:"address" mapstructure:"address"`
	CreationTimestamp int64  `json:"creation_timestamp" mapstructure:"creation_timestamp"`
	Currency          string `json:"currency" mapstructure:"currency"`
	Type              string `json:"type" mapstructure:"type"`
}
