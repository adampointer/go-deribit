package public

type GetAnnouncementsResponse []struct {
	Body            string `json:"body" mapstructure:"body"`
	ID              int64  `json:"id" mapstructure:"id"`
	Important       bool   `json:"important" mapstructure:"important"`
	PublicationTime int64  `json:"publication_time" mapstructure:"publication_time"`
	Title           string `json:"title" mapstructure:"title"`
}
