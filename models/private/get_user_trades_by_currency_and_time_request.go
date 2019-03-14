package private

type GetUserTradesByCurrencyAndTimeRequest struct {
	Count          int64  `json:"count"`
	Currency       string `json:"currency"`
	EndTimestamp   int64  `json:"end_timestamp"`
	StartTimestamp int64  `json:"start_timestamp"`
}
