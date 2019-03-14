package private

type SetPasswordForSubaccountRequest struct {
	Password string `json:"password"`
	Sid      int64  `json:"sid"`
}
