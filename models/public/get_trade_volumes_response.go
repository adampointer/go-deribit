package public

type GetTradeVolumesResponse []struct {
	CallsVolume   int64   `json:"calls_volume" mapstructure:"calls_volume"`
	CurrencyPair  string  `json:"currency_pair" mapstructure:"currency_pair"`
	FuturesVolume float64 `json:"futures_volume" mapstructure:"futures_volume"`
	PutsVolume    int64   `json:"puts_volume" mapstructure:"puts_volume"`
}
