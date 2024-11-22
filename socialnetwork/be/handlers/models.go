package handlers

import "encoding/json"

type UserProfileRequest struct {
	ExternalID         int64           `json:"external_id"`
	Name               string          `json:"name"`
	Info               string          `json:"info"`
	Age                int             `json:"age"`
	Occupation         string          `json:"occupation"`
	Personality        json.RawMessage `json:"personality"`
	Interests          json.RawMessage `json:"interests"`
	Background         json.RawMessage `json:"background"`
	CommunicationStyle json.RawMessage `json:"communication_style"`
	SocialConnections  json.RawMessage `json:"social_connections"`
}

type UserFullProfileResponse struct {
	ID                 int                `json:"id"`
	ExternalID         int64              `json:"external_id"`
	Name               string             `json:"name"`
	Info               string             `json:"info"`
	Age                int                `json:"age"`
	Occupation         string             `json:"occupation"`
	Personality        Personality        `json:"personality"`
	Interests          Interests          `json:"interests"`
	Background         Background         `json:"background"`
	CommunicationStyle CommunicationStyle `json:"communication_style"`
	SocialConnections  SocialConnections  `json:"social_connections"`
}

type Personality struct {
	Type             string `json:"type"`
	HumorStyle       string `json:"humor_style"`
	SocialPreference string `json:"social_preference"`
}

type Interests struct {
	PrimaryInterests   []string `json:"primary_interests"`
	SecondaryInterests []string `json:"secondary_interests"`
	PreferredTopics    []string `json:"preferred_topics"`
	DislikedTopics     []string `json:"disliked_topics"`
}

type Background struct {
	Hometown       string   `json:"hometown"`
	EducationLevel string   `json:"education_level"`
	Values         []string `json:"values"`
}

type CommunicationStyle struct {
	FavoriteWords        []string `json:"favorite_words"`
	FormalityLevel       string   `json:"formality_level"`
	ConversationTendency string   `json:"conversation_tendency"`
}

type SocialConnections struct {
	Groups             []string `json:"groups"`
	FriendlinessRating string   `json:"friendliness_rating"`
}
