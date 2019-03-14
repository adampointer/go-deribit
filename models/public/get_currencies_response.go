package public

type GetCurrenciesResponse []struct {
	CoinType                       string  `json:"coin_type"`
	Currency                       string  `json:"currency"`
	CurrencyLong                   string  `json:"currency_long"`
	DisabledDepositAddressCreation bool    `json:"disabled_deposit_address_creation"`
	FeePrecision                   int64   `json:"fee_precision"`
	MinConfirmations               int64   `json:"min_confirmations"`
	MinWithdrawalFee               float64 `json:"min_withdrawal_fee"`
	WithdrawalFee                  float64 `json:"withdrawal_fee"`
	WithdrawalPriorities           []struct {
		Name  string  `json:"name"`
		Value float64 `json:"value"`
	} `json:"withdrawal_priorities"`
}
