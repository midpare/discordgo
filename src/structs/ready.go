package structs

type Ready struct {
	Version          int      `json:"v"`
	User             *User    `json:"user"`
	Guilds           []*Guild `json:"guilds"`
	SessionId        string   `json:"session_id"`
	ResumeGatewayUrl string   `json:"resume_gateway_url"`
}
