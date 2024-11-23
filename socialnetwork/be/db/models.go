package db

type NewUserFullProfile struct {
	ExternalID         int64
	Name               string
	Info               string
	Age                int
	Occupation         string
	Personality        Personality
	Interests          Interests
	Background         Background
	CommunicationStyle CommunicationStyle
	SocialConnections  SocialConnections
}

type UserFullProfile struct {
	ID                 int
	ExternalID         int64
	Name               string
	Info               string
	Age                int
	Occupation         string
	Personality        Personality
	Interests          Interests
	Background         Background
	CommunicationStyle CommunicationStyle
	SocialConnections  SocialConnections
}

type Personality struct {
	Type             string
	HumorStyle       string
	SocialPreference string
}

type Interests struct {
	PrimaryInterests   []string
	SecondaryInterests []string
	PreferredTopics    []string
	DislikedTopics     []string
}

type Background struct {
	Hometown       string
	EducationLevel string
	Values         []string
}

type CommunicationStyle struct {
	FavoriteWords        []string
	FormalityLevel       string
	ConversationTendency string
}

type SocialConnections struct {
	Groups             []string
	FriendlinessRating string
}
