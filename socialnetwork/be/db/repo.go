package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type UserRepo struct {
	DB *sql.DB
}

type UserRepository interface {
	QueryProfiles() ([]UserProfile, error)
	CreateProfile(profileRequest UserProfileRequest) (UserProfile, error)
}

const FULL_PROFILE_QUERY = `
    SELECT *
    FROM user_profiles
`

func (u UserRepo) QueryProfiles() ([]UserProfile, error) {
	rows, err := u.DB.Query(FULL_PROFILE_QUERY)

	if err != nil {
		return nil, fmt.Errorf("failed to query full user profiles: %w", err)
	}
	defer rows.Close()

	var profiles []UserProfile = make([]UserProfile, 0)
	for rows.Next() {
		var profile UserProfile

		err := rows.Scan(
			&profile.ID,
			&profile.Traits,
			&profile.Profile,
		)
		fmt.Print("LOADED PROFILE", profile.Profile)
		if err != nil {
			return nil, fmt.Errorf("Error scanning row: %v", err)
		}
		profiles = append(profiles, profile)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error iterating over rows: %v", err)
	}
	return profiles, nil
}

func (u UserRepo) CreateProfile(profileRequest UserProfileRequest) (UserProfile, error) {
	tx, err := u.DB.Begin()
	if err != nil {
		return UserProfile{}, fmt.Errorf("Error starting transaction: %v", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	fmt.Println("SAVING ", profileRequest)

	var userProfileID int
	err = tx.QueryRow(
		"INSERT INTO user_profiles (traits, profile) VALUES ($1, $2) RETURNING id",
		profileRequest.Traits,
		profileRequest.Profile,
	).Scan(&userProfileID)
	if err != nil {
		return UserProfile{}, fmt.Errorf("Error inserting into user_profile: %v", err)
	}

	if err := tx.Commit(); err != nil {
		return UserProfile{}, fmt.Errorf("Error committing transaction: %v", err)
	}

	profileResponse := UserProfile{
		ID:      userProfileID,
		Traits:  profileRequest.Traits,
		Profile: profileRequest.Profile,
	}

	return profileResponse, nil
}
