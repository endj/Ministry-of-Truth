package db

type MockUserRepository struct {
}

func (m *MockUserRepository) QueryFullProfiles() ([]UserFullProfile, error) {
	return []UserFullProfile{
		{
			ID:         1,
			ExternalID: 12345,
			Name:       "John Doe",
			Info:       "Software Engineer",
			Age:        30,
			Occupation: "Engineer",
			Personality: Personality{
				Type:             "Introvert",
				HumorStyle:       "Dry",
				SocialPreference: "Reserved",
			},
			Interests: Interests{
				PrimaryInterests:   []string{"Coding", "Music"},
				SecondaryInterests: []string{"Hiking"},
				PreferredTopics:    []string{"Tech", "Science"},
				DislikedTopics:     []string{"Politics"},
			},
			Background: Background{
				Hometown:       "New York",
				EducationLevel: "Bachelors",
				Values:         []string{"Integrity", "Innovation"},
			},
			CommunicationStyle: CommunicationStyle{
				FavoriteWords:        []string{"Efficiency", "Clarity"},
				FormalityLevel:       "Formal",
				ConversationTendency: "Structured",
			},
			SocialConnections: SocialConnections{
				Groups:             []string{"Work", "Family"},
				FriendlinessRating: "High",
			},
		},
	}, nil
}

func (m *MockUserRepository) InsertUserProfileWithTemplate(profile NewUserFullProfile) (UserFullProfile, error) {
	return UserFullProfile{
		ID:         1,
		ExternalID: profile.ExternalID,
		Name:       profile.Name,
		Info:       profile.Info,
		Age:        profile.Age,
		Occupation: profile.Occupation,
		Personality: Personality{
			Type:             profile.Personality.Type,
			HumorStyle:       profile.Personality.HumorStyle,
			SocialPreference: profile.Personality.SocialPreference,
		},
		Interests: Interests{
			PrimaryInterests:   profile.Interests.PrimaryInterests,
			SecondaryInterests: profile.Interests.SecondaryInterests,
			PreferredTopics:    profile.Interests.PreferredTopics,
			DislikedTopics:     profile.Interests.DislikedTopics,
		},
		Background: Background{
			Hometown:       profile.Background.Hometown,
			EducationLevel: profile.Background.EducationLevel,
			Values:         profile.Background.Values,
		},
		CommunicationStyle: CommunicationStyle{
			FavoriteWords:        profile.CommunicationStyle.FavoriteWords,
			FormalityLevel:       profile.CommunicationStyle.FormalityLevel,
			ConversationTendency: profile.CommunicationStyle.ConversationTendency,
		},
		SocialConnections: SocialConnections{
			Groups:             profile.SocialConnections.Groups,
			FriendlinessRating: profile.SocialConnections.FriendlinessRating,
		},
	}, nil
}
