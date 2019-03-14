package public

type GetBookSummaryByCurrencyResponse []struct {
	AskPrice               interface{} `json:"ask_price" mapstructure:"ask_price"`
	BaseCurrency           string      `json:"base_currency" mapstructure:"base_currency"`
	BidPrice               interface{} `json:"bid_price" mapstructure:"bid_price"`
	CreationTimestamp      int64       `json:"creation_timestamp" mapstructure:"creation_timestamp"`
	CurrentFunding         int64       `json:"current_funding" mapstructure:"current_funding"`
	EstimatedDeliveryPrice float64     `json:"estimated_delivery_price" mapstructure:"estimated_delivery_price"`
	Funding8h              float64     `json:"funding_8h" mapstructure:"funding_8h"`
	High                   int64       `json:"high" mapstructure:"high"`
	InstrumentName         string      `json:"instrument_name" mapstructure:"instrument_name"`
	Last                   int64       `json:"last" mapstructure:"last"`
	Low                    int64       `json:"low" mapstructure:"low"`
	MarkPrice              float64     `json:"mark_price" mapstructure:"mark_price"`
	MidPrice               interface{} `json:"mid_price" mapstructure:"mid_price"`
	OpenInterest           int64       `json:"open_interest" mapstructure:"open_interest"`
	QuoteCurrency          string      `json:"quote_currency" mapstructure:"quote_currency"`
	Volume                 int64       `json:"volume" mapstructure:"volume"`
	VolumeUsd              int64       `json:"volume_usd" mapstructure:"volume_usd"`
}
