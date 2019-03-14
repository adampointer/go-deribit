package private

type GetSubaccountsResponse []struct {
	Email        string `json:"email"`
	ID           int64  `json:"id"`
	IsPassword   bool   `json:"is_password"`
	LoginEnabled bool   `json:"login_enabled"`
	Portfolio    struct {
		Btc struct {
			AvailableFunds           float64 `json:"available_funds"`
			AvailableWithdrawalFunds float64 `json:"available_withdrawal_funds"`
			Balance                  float64 `json:"balance"`
			Currency                 string  `json:"currency"`
			Equity                   float64 `json:"equity"`
			InitialMargin            float64 `json:"initial_margin"`
			MaintenanceMargin        float64 `json:"maintenance_margin"`
			MarginBalance            float64 `json:"margin_balance"`
		} `json:"btc"`
		Eth struct {
			AvailableFunds           int64  `json:"available_funds"`
			AvailableWithdrawalFunds int64  `json:"available_withdrawal_funds"`
			Balance                  int64  `json:"balance"`
			Currency                 string `json:"currency"`
			Equity                   int64  `json:"equity"`
			InitialMargin            int64  `json:"initial_margin"`
			MaintenanceMargin        int64  `json:"maintenance_margin"`
			MarginBalance            int64  `json:"margin_balance"`
		} `json:"eth"`
	} `json:"portfolio"`
	ReceiveNotifications bool   `json:"receive_notifications"`
	SystemName           string `json:"system_name"`
	TfaEnabled           bool   `json:"tfa_enabled"`
	Type                 string `json:"type"`
	Username             string `json:"username"`
}
