package private

type GetTransfersResponse struct {
	Count int64 `json:"count" mapstructure:"count"`
	Data  []struct {
		Amount           float64 `json:"amount" mapstructure:"amount"`
		CreatedTimestamp int64   `json:"created_timestamp" mapstructure:"created_timestamp"`
		Currency         string  `json:"currency" mapstructure:"currency"`
		Direction        string  `json:"direction" mapstructure:"direction"`
		ID               int64   `json:"id" mapstructure:"id"`
		OtherSide        string  `json:"other_side" mapstructure:"other_side"`
		State            string  `json:"state" mapstructure:"state"`
		Type             string  `json:"type" mapstructure:"type"`
		UpdatedTimestamp int64   `json:"updated_timestamp" mapstructure:"updated_timestamp"`
	} `json:"data" mapstructure:"data"`
}
