package private

type CancelTransferByIdRequest struct {
	Currency string `json:"currency"`
	ID       int64  `json:"id"`
}
