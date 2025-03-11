package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type UserRepo struct {
	DB *sql.DB
}

func (u UserRepo) QueryProfiles() ([]UserProfile, error) {
	rows, err := u.DB.Query(`
    SELECT *
    FROM user_profiles
    `)
	if err != nil {
		return nil, fmt.Errorf("failed to query full user profiles: %w", err)
	}

	defer rows.Close()

	var profiles []UserProfile = make([]UserProfile, 0)
	for rows.Next() {
		var profile UserProfile
		if err := rows.Scan(&profile.ID, &profile.Name, &profile.Traits, &profile.Profile); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		profiles = append(profiles, profile)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}
	return profiles, nil
}

func (u UserRepo) CreateProfile(profileRequest UserProfileRequest) (*UserProfile, error) {
	tx, err := u.DB.Begin()
	if err != nil {
		return nil, fmt.Errorf("error starting transaction: %v", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	log.Println("Creating profile: ", profileRequest)

	var userProfileID int
	if err := tx.QueryRow(
		"INSERT INTO user_profiles (name, traits, profile) VALUES ($1, $2, $3) RETURNING id",
		profileRequest.Name, profileRequest.Traits, profileRequest.Profile).Scan(&userProfileID); err != nil {
		return nil, fmt.Errorf("error inserting into user_profile: %v", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("error committing transaction: %v", err)
	}

	return &UserProfile{
		ID:      userProfileID,
		Name:    profileRequest.Name,
		Traits:  profileRequest.Traits,
		Profile: profileRequest.Profile,
	}, nil
}
