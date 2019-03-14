package private

type SetEmailForSubaccountRequest struct {
	Email string `json:"email"`
	Sid   int64  `json:"sid"`
}
