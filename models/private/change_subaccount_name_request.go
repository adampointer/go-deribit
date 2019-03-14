package private

type ChangeSubaccountNameRequest struct {
	Name string `json:"name" mapstructure:"name"`
	Sid  int64  `json:"sid" mapstructure:"sid"`
}
