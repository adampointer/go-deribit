package public

type HelloRequest struct {
	ClientName    string `json:"client_name"`
	ClientVersion string `json:"client_version"`
}
