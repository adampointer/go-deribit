package public

type GetTradeVolumesResponse []struct {
	CallsVolume   int64   `json:"calls_volume"`
	CurrencyPair  string  `json:"currency_pair"`
	FuturesVolume float64 `json:"futures_volume"`
	PutsVolume    int64   `json:"puts_volume"`
}
