package private

type SetEmailForSubaccountRequest struct {
	Email string `json:"email" mapstructure:"email"`
	Sid   int64  `json:"sid" mapstructure:"sid"`
}
