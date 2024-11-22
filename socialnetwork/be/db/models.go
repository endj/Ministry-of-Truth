package db

type UserFullProfile struct {
	ID                 int    `json:"id"`
	ExternalID         int64  `json:"external_id"`
	Name               string `json:"name"`
	Info               string `json:"info"`
	Age                int    `json:"age"`
	Occupation         string `json:"occupation"`
	Personality        []byte `json:"personality"`
	Interests          []byte `json:"interests"`
	Background         []byte `json:"background"`
	CommunicationStyle []byte `json:"communication_style"`
	SocialConnections  []byte `json:"social_connections"` // JSONB as byte array
}
