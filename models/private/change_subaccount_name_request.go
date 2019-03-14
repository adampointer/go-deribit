package private

type ChangeSubaccountNameRequest struct {
	Name string `json:"name"`
	Sid  int64  `json:"sid"`
}
