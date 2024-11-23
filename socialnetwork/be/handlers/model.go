package handlers

type UserFullProfileRequest struct {
	ExternalID            int64    `json:"external_id"`
	Name                  string   `json:"name"`
	Info                  string   `json:"info"`
	Age                   int      `json:"age"`
	Occupation            string   `json:"occupation"`
	PersonalityType       string   `json:"personality_type"`
	PersonalityHumor      string   `json:"personality_humor_style"`
	PersonalitySocial     string   `json:"personality_social_preference"`
	InterestsPrimary      []string `json:"interests_primary_interests"`
	InterestsSecondary    []string `json:"interests_secondary_interests"`
	InterestsPreferred    []string `json:"interests_preferred_topics"`
	InterestsDisliked     []string `json:"interests_disliked_topics"`
	BackgroundHometown    string   `json:"background_hometown"`
	BackgroundEducation   string   `json:"background_education_level"`
	BackgroundValues      []string `json:"background_values"`
	CommStyleFavorites    []string `json:"communication_style_favorite_words"`
	CommStyleFormality    string   `json:"communication_style_formality_level"`
	CommStyleConversation string   `json:"communication_style_conversation_tendency"`
	SocialGroups          []string `json:"social_connections_groups"`
	SocialFriendliness    string   `json:"social_connections_friendliness_rating"`
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
