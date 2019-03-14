package private

type ClosePositionRequest struct {
	InstrumentName string  `json:"instrument_name"`
	Price          float64 `json:"price"`
	Type           string  `json:"type"`
}
