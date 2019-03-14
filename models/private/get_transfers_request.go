package private

type GetTransfersRequest struct {
	Count    int64  `json:"count"`
	Currency string `json:"currency"`
}
