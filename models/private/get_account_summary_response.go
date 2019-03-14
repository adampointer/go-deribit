package private

type GetAccountSummaryResponse struct {
	AvailableFunds            float64 `json:"available_funds"`
	AvailableWithdrawalFunds  float64 `json:"available_withdrawal_funds"`
	Balance                   float64 `json:"balance"`
	Currency                  string  `json:"currency"`
	DeltaTotal                float64 `json:"delta_total"`
	DepositAddress            string  `json:"deposit_address"`
	Email                     string  `json:"email"`
	Equity                    float64 `json:"equity"`
	FuturesPl                 int64   `json:"futures_pl"`
	FuturesSessionRpl         int64   `json:"futures_session_rpl"`
	FuturesSessionUpl         int64   `json:"futures_session_upl"`
	ID                        int64   `json:"id"`
	InitialMargin             float64 `json:"initial_margin"`
	MaintenanceMargin         float64 `json:"maintenance_margin"`
	MarginBalance             float64 `json:"margin_balance"`
	OptionsDelta              int64   `json:"options_delta"`
	OptionsGamma              int64   `json:"options_gamma"`
	OptionsPl                 int64   `json:"options_pl"`
	OptionsSessionRpl         int64   `json:"options_session_rpl"`
	OptionsSessionUpl         int64   `json:"options_session_upl"`
	OptionsTheta              int64   `json:"options_theta"`
	OptionsVega               int64   `json:"options_vega"`
	PortfolioMarginingEnabled bool    `json:"portfolio_margining_enabled"`
	SessionFunding            int64   `json:"session_funding"`
	SessionRpl                int64   `json:"session_rpl"`
	SessionUpl                int64   `json:"session_upl"`
	SystemName                string  `json:"system_name"`
	TfaEnabled                bool    `json:"tfa_enabled"`
	TotalPl                   int64   `json:"total_pl"`
	Type                      string  `json:"type"`
	Username                  string  `json:"username"`
}
