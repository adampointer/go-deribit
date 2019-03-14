package private

type GetNewAnnouncementsResponse []struct {
	Body            string `json:"body"`
	ID              int64  `json:"id"`
	Important       bool   `json:"important"`
	PublicationTime int64  `json:"publication_time"`
	Title           string `json:"title"`
}
