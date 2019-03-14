package notifications

type UserPortfolioCurrencyRepeated struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  struct {
		Channel string `json:"channel"`
		Data    struct {
			AvailableFunds            float64 `json:"available_funds"`
			AvailableWithdrawalFunds  float64 `json:"available_withdrawal_funds"`
			Balance                   float64 `json:"balance"`
			Currency                  string  `json:"currency"`
			DeltaTotal                float64 `json:"delta_total"`
			Equity                    float64 `json:"equity"`
			FuturesPl                 float64 `json:"futures_pl"`
			FuturesSessionRpl         int64   `json:"futures_session_rpl"`
			FuturesSessionUpl         float64 `json:"futures_session_upl"`
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
			SessionUpl                float64 `json:"session_upl"`
			TotalPl                   float64 `json:"total_pl"`
		} `json:"data"`
	} `json:"params"`
}
