package private

type ToggleSubaccountLoginRequest struct {
	Sid   int64  `json:"sid" mapstructure:"sid"`
	State string `json:"state" mapstructure:"state"`
}
