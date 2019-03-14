package public

type GetBookSummaryByCurrencyResponse []struct {
	AskPrice               interface{} `json:"ask_price"`
	BaseCurrency           string      `json:"base_currency"`
	BidPrice               interface{} `json:"bid_price"`
	CreationTimestamp      int64       `json:"creation_timestamp"`
	CurrentFunding         int64       `json:"current_funding"`
	EstimatedDeliveryPrice float64     `json:"estimated_delivery_price"`
	Funding8h              float64     `json:"funding_8h"`
	High                   int64       `json:"high"`
	InstrumentName         string      `json:"instrument_name"`
	Last                   int64       `json:"last"`
	Low                    int64       `json:"low"`
	MarkPrice              float64     `json:"mark_price"`
	MidPrice               interface{} `json:"mid_price"`
	OpenInterest           int64       `json:"open_interest"`
	QuoteCurrency          string      `json:"quote_currency"`
	Volume                 int64       `json:"volume"`
	VolumeUsd              int64       `json:"volume_usd"`
}
