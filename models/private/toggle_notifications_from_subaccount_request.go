package private

type ToggleNotificationsFromSubaccountRequest struct {
	Sid   int64 `json:"sid"`
	State bool  `json:"state"`
}
