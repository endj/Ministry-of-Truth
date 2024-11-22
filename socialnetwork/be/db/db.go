package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitializeDB(dataSourceName string) {
	var err error
	DB, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	if err := DB.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}
}

func QueryFullProfiles() (*sql.Rows, error) {
	query := `
		SELECT
			up.id,
			up.external_id,
			up.name,
			up.info,
			ut.age,
			ut.occupation,
			ut.personality,
			ut.interests,
			ut.background,
			ut.communication_style,
			ut.social_connections
		FROM user_profiles up
		LEFT JOIN user_templates ut ON ut.user_profile_id = up.id
	`
	return DB.Query(query)
}

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

func InsertUserProfileWithTemplate(profileRequest UserProfileRequest) (UserFullProfile, error) {
	tx, err := DB.Begin()
	if err != nil {
		return UserFullProfile{}, fmt.Errorf("Error starting transaction: %v", err)
	}

	defer tx.Rollback()

	var userProfileID int
	err = tx.QueryRow(
		"INSERT INTO user_profile (external_id, name, info) VALUES ($1, $2, $3) RETURNING id",
		profileRequest.ExternalID, profileRequest.Name, profileRequest.Info,
	).Scan(&userProfileID)
	if err != nil {
		return UserFullProfile{}, fmt.Errorf("Error inserting into user_profile: %v", err)
	}

	_, err = tx.Exec(
		"INSERT INTO user_templates (user_profile_id, name, age, occupation, personality, interests, background, communication_style, social_connections) "+
			"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		userProfileID,
		profileRequest.Name, profileRequest.Age, profileRequest.Occupation,
		profileRequest.Personality, profileRequest.Interests, profileRequest.Background,
		profileRequest.CommunicationStyle, profileRequest.SocialConnections,
	)
	if err != nil {
		return UserFullProfile{}, fmt.Errorf("Error inserting into user_templates: %v", err)
	}

	if err := tx.Commit(); err != nil {
		return UserFullProfile{}, fmt.Errorf("Error committing transaction: %v", err)
	}

	profileResponse := UserFullProfile{
		ID:                 userProfileID,
		ExternalID:         profileRequest.ExternalID,
		Name:               profileRequest.Name,
		Info:               profileRequest.Info,
		Age:                profileRequest.Age,
		Occupation:         profileRequest.Occupation,
		Personality:        profileRequest.Personality,
		Interests:          profileRequest.Interests,
		Background:         profileRequest.Background,
		CommunicationStyle: profileRequest.CommunicationStyle,
		SocialConnections:  profileRequest.SocialConnections,
	}

	return profileResponse, nil
}
