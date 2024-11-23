package db

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/lib/pq"
)

type UserRepo struct {
	DB *sql.DB
}

type UserRepository interface {
	QueryFullProfiles() ([]UserFullProfile, error)
	InsertUserProfileWithTemplate(profileRequest NewUserFullProfile) (UserFullProfile, error)
}

const FULL_PROFILE_QUERY = `
    SELECT
        up.id,
        up.external_id,
        up.name,
        up.info,
        ut.template
    FROM user_profiles up
    LEFT JOIN user_templates ut ON ut.user_profile_id = up.id
`

func (u UserRepo) QueryFullProfiles() ([]UserFullProfile, error) {
	rows, err := u.DB.Query(FULL_PROFILE_QUERY)

	if err != nil {
		return nil, fmt.Errorf("failed to query full user profiles: %w", err)
	}
	defer rows.Close()

	var profiles []UserFullProfile
	for rows.Next() {
		var profile UserFullProfile
		var template []byte

		// Scan the columns: profile data from user_profiles and the JSONB template from user_templates
		err := rows.Scan(
			&profile.ID,
			&profile.ExternalID,
			&profile.Name,
			&profile.Info,
			&template, // Get the template field
		)
		if err != nil {
			return nil, fmt.Errorf("Error scanning row: %v", err)
		}

		// Unmarshal the JSONB template field into the UserFullProfile struct
		if err := json.Unmarshal(template, &profile); err != nil {
			return nil, fmt.Errorf("Error unmarshaling template: %v", err)
		}

		profiles = append(profiles, profile)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error iterating over rows: %v", err)
	}
	return profiles, nil
}

func (u UserRepo) InsertUserProfileWithTemplate(profileRequest NewUserFullProfile) (UserFullProfile, error) {
	tx, err := u.DB.Begin()
	if err != nil {
		return UserFullProfile{}, fmt.Errorf("Error starting transaction: %v", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	var userProfileID int
	err = tx.QueryRow(
		"INSERT INTO user_profiles (external_id, name, info) VALUES ($1, $2, $3) RETURNING id",
		profileRequest.ExternalID, profileRequest.Name, profileRequest.Info,
	).Scan(&userProfileID)
	if err != nil {
		return UserFullProfile{}, fmt.Errorf("Error inserting into user_profile: %v", err)
	}

	templateJSON, err := json.Marshal(profileRequest)
	if err != nil {
		return UserFullProfile{}, fmt.Errorf("Error marshalling profile request to JSON: %v", err)
	}

	_, err = tx.Exec(
		"INSERT INTO user_templates (user_profile_id, template) VALUES ($1, $2)",
		userProfileID, templateJSON,
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
