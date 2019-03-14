package private

type GetMarginsRequest struct {
	Amount         int64  `json:"amount"`
	InstrumentName string `json:"instrument_name"`
	Price          int64  `json:"price"`
}
