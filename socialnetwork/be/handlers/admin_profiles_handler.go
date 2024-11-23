package handlers

import (
	"app/db"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func AdmingListUsersHandler(repo db.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		setHeaders(w)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		profiles, err := repo.QueryFullProfiles()
		if err != nil {
			http.Error(w, fmt.Sprintf("Error querying database: %v", err), http.StatusInternalServerError)
			return
		}

		var responseProfiles []UserFullProfileResponse
		for _, profile := range profiles {
			responseProfiles = append(responseProfiles, fromFullProfileToResponse(profile))
		}

		if err := json.NewEncoder(w).Encode(responseProfiles); err != nil {
			http.Error(w, fmt.Sprintf("Error encoding response: %v", err), http.StatusInternalServerError)
			return
		}
	}
}

func AdminCreateProfileHandler(repo db.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		setHeaders(w)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		var profileRequest UserFullProfileRequest
		if err := json.NewDecoder(r.Body).Decode(&profileRequest); err != nil {

			http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
			return
		}

		if err := profileRequest.Validate(); err != nil {
			http.Error(w, fmt.Sprintf("Validation failed: %v", err), http.StatusBadRequest)
			return
		}

		profileResponse, err := repo.InsertUserProfileWithTemplate(toNewUserFullProfile(profileRequest))
		if err != nil {
			http.Error(w, fmt.Sprintf("Error processing request: %v", err), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(profileResponse)
	}
}

func setHeaders(w http.ResponseWriter) {
	// Allow CORS for all origins
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	// Set content type for the response
	w.Header().Set("Content-Type", "application/json")

	// If it's an OPTIONS request, respond with a 200 status code and exit

}

func fromFullProfileToResponse(profile db.UserFullProfile) UserFullProfileResponse {
	return UserFullProfileResponse{
		ID:         profile.ID,
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
	}
}

func toNewUserFullProfile(req UserFullProfileRequest) db.NewUserFullProfile {
	personality := db.Personality{
		Type:             req.PersonalityType,
		HumorStyle:       req.PersonalityHumor,
		SocialPreference: req.PersonalitySocial,
	}

	interests := db.Interests{
		PrimaryInterests:   req.InterestsPrimary,
		SecondaryInterests: req.InterestsSecondary,
		PreferredTopics:    req.InterestsPreferred,
		DislikedTopics:     req.InterestsDisliked,
	}

	background := db.Background{
		Hometown:       req.BackgroundHometown,
		EducationLevel: req.BackgroundEducation,
		Values:         req.BackgroundValues,
	}

	communicationStyle := db.CommunicationStyle{
		FavoriteWords:        req.CommStyleFavorites,
		FormalityLevel:       req.CommStyleFormality,
		ConversationTendency: req.CommStyleConversation,
	}

	socialConnections := db.SocialConnections{
		Groups:             req.SocialGroups,
		FriendlinessRating: req.SocialFriendliness,
	}

	return db.NewUserFullProfile{
		ExternalID:         req.ExternalID,
		Name:               req.Name,
		Info:               req.Info,
		Age:                req.Age,
		Occupation:         req.Occupation,
		Personality:        personality,
		Interests:          interests,
		Background:         background,
		CommunicationStyle: communicationStyle,
		SocialConnections:  socialConnections,
	}
}

func (u *UserFullProfileRequest) Validate() error {
	// Check if required fields are empty
	if u.ExternalID == 0 {
		return fmt.Errorf("external_id is required")
	}
	if strings.TrimSpace(u.Name) == "" {
		return fmt.Errorf("name is required")
	}
	if strings.TrimSpace(u.Info) == "" {
		return fmt.Errorf("info is required")
	}
	if u.Age <= 0 {
		return fmt.Errorf("age is required and must be a positive number")
	}
	if strings.TrimSpace(u.Occupation) == "" {
		return fmt.Errorf("occupation is required")
	}
	if strings.TrimSpace(u.PersonalityType) == "" {
		return fmt.Errorf("personality_type is required")
	}
	if strings.TrimSpace(u.PersonalityHumor) == "" {
		return fmt.Errorf("personality_humor_style is required")
	}
	if strings.TrimSpace(u.PersonalitySocial) == "" {
		return fmt.Errorf("personality_social_preference is required")
	}
	if len(u.InterestsPrimary) == 0 {
		return fmt.Errorf("interests_primary_interests is required")
	}
	if len(u.InterestsSecondary) == 0 {
		return fmt.Errorf("interests_secondary_interests is required")
	}
	if len(u.InterestsPreferred) == 0 {
		return fmt.Errorf("interests_preferred_topics is required")
	}
	if len(u.InterestsDisliked) == 0 {
		return fmt.Errorf("interests_disliked_topics is required")
	}
	if strings.TrimSpace(u.BackgroundHometown) == "" {
		return fmt.Errorf("background_hometown is required")
	}
	if strings.TrimSpace(u.BackgroundEducation) == "" {
		return fmt.Errorf("background_education_level is required")
	}
	if len(u.BackgroundValues) == 0 {
		return fmt.Errorf("background_values is required")
	}
	if len(u.CommStyleFavorites) == 0 {
		return fmt.Errorf("communication_style_favorite_words is required")
	}
	if strings.TrimSpace(u.CommStyleFormality) == "" {
		return fmt.Errorf("communication_style_formality_level is required")
	}
	if strings.TrimSpace(u.CommStyleConversation) == "" {
		return fmt.Errorf("communication_style_conversation_tendency is required")
	}
	if len(u.SocialGroups) == 0 {
		return fmt.Errorf("social_connections_groups is required")
	}
	if strings.TrimSpace(u.SocialFriendliness) == "" {
		return fmt.Errorf("social_connections_friendliness_rating is required")
	}

	// If all validations passed, return nil
	return nil
}
