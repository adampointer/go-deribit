package private

type CancelAllByInstrumentRequest struct {
	InstrumentName string `json:"instrument_name"`
	Type           string `json:"type"`
}
