package public

type GetInstrumentsResponse []struct {
	BaseCurrency        string  `json:"base_currency"`
	ContractSize        int64   `json:"contract_size"`
	CreationTimestamp   int64   `json:"creation_timestamp"`
	ExpirationTimestamp int64   `json:"expiration_timestamp"`
	InstrumentName      string  `json:"instrument_name"`
	IsActive            bool    `json:"is_active"`
	Kind                string  `json:"kind"`
	MinTradeAmount      int64   `json:"min_trade_amount"`
	QuoteCurrency       string  `json:"quote_currency"`
	SettlementPeriod    string  `json:"settlement_period"`
	TickSize            float64 `json:"tick_size"`
}
