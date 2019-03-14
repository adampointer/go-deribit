package public

type GetLastSettlementsByInstrumentResponse struct {
	Continuation string `json:"continuation"`
	Settlements  []struct {
		IndexPrice        float64 `json:"index_price"`
		InstrumentName    string  `json:"instrument_name"`
		MarkPrice         float64 `json:"mark_price"`
		Position          int64   `json:"position"`
		ProfitLoss        int64   `json:"profit_loss"`
		SessionProfitLoss float64 `json:"session_profit_loss"`
		Timestamp         int64   `json:"timestamp"`
		Type              string  `json:"type"`
	} `json:"settlements"`
}
