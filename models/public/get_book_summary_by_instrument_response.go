package public

type GetBookSummaryByInstrumentResponse []struct {
	AskPrice          float64 `json:"ask_price"`
	BaseCurrency      string  `json:"base_currency"`
	BidPrice          float64 `json:"bid_price"`
	CreationTimestamp int64   `json:"creation_timestamp"`
	High              float64 `json:"high"`
	InstrumentName    string  `json:"instrument_name"`
	InterestRate      float64 `json:"interest_rate"`
	Last              float64 `json:"last"`
	Low               float64 `json:"low"`
	MarkPrice         float64 `json:"mark_price"`
	MidPrice          float64 `json:"mid_price"`
	OpenInterest      float64 `json:"open_interest"`
	QuoteCurrency     string  `json:"quote_currency"`
	UnderlyingIndex   string  `json:"underlying_index"`
	UnderlyingPrice   float64 `json:"underlying_price"`
	Volume            float64 `json:"volume"`
}
