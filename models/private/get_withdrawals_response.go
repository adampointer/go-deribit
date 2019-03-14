package private

type GetWithdrawalsResponse struct {
	Count int64 `json:"count"`
	Data  []struct {
		Address            string      `json:"address"`
		Amount             float64     `json:"amount"`
		ConfirmedTimestamp interface{} `json:"confirmed_timestamp"`
		CreatedTimestamp   int64       `json:"created_timestamp"`
		Currency           string      `json:"currency"`
		Fee                float64     `json:"fee"`
		ID                 int64       `json:"id"`
		Priority           float64     `json:"priority"`
		State              string      `json:"state"`
		TransactionID      interface{} `json:"transaction_id"`
		UpdatedTimestamp   int64       `json:"updated_timestamp"`
	} `json:"data"`
}
