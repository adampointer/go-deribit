package private

type GetCurrentDepositAddressResponse struct {
	Address           string `json:"address"`
	CreationTimestamp int64  `json:"creation_timestamp"`
	Currency          string `json:"currency"`
	Type              string `json:"type"`
}
