package private

type SetPasswordForSubaccountRequest struct {
	Password string `json:"password" mapstructure:"password"`
	Sid      int64  `json:"sid" mapstructure:"sid"`
}
