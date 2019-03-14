package private

type ToggleSubaccountLoginRequest struct {
	Sid   int64  `json:"sid"`
	State string `json:"state"`
}
