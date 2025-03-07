package handlers

import (
	"app/db"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetProfilesHandler(w http.ResponseWriter, r *http.Request, repo db.UserRepo) {
	profiles, err := repo.QueryProfiles()
	if err != nil {
		http.Error(w, fmt.Sprintf("error querying database: %v", err), http.StatusInternalServerError)
		return
	}

	var responseProfiles []UserProfile = make([]UserProfile, 0)
	for _, profile := range profiles {
		responseProfiles = append(responseProfiles, toProfileResponse(&profile))
	}

	if err := json.NewEncoder(w).Encode(responseProfiles); err != nil {
		http.Error(w, fmt.Sprintf("error encoding response: %v", err), http.StatusInternalServerError)
	}
}

func CreateProfileHandler(w http.ResponseWriter, r *http.Request, repo db.UserRepo) {
	request, err := toProfileRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	profileResponse, err := repo.CreateProfile(*request)
	if err != nil {
		http.Error(w, fmt.Sprintf("error processing request: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(toProfileResponse(profileResponse)); err != nil {
		http.Error(w, fmt.Sprintf("error encoding response: %v", err), http.StatusInternalServerError)
	}
}

func toJsonString(m map[string]any) (string, error) {
	bytes, err := json.Marshal(m)
	if err != nil {
		return "", fmt.Errorf("failed to marshal map to JSON: %w", err)
	}
	return string(bytes), nil
}

func toProfileRequest(r *http.Request) (*db.UserProfileRequest, error) {
	var profileRequest UserProfileRequest
	var traits string
	var profile string

	if err := json.NewDecoder(r.Body).Decode(&profileRequest); err != nil {
		log.Println("Failed to decode body")
		return nil, fmt.Errorf("invalid JSON payload: " + err.Error())
	}
	if profileRequest.Profile == nil || profileRequest.Traits == nil {
		log.Println("Profile or traits missing")
		return nil, fmt.Errorf("invalid JSON payload: missing fields profile or traits")
	}
	var err error
	if traits, err = toJsonString(profileRequest.Traits); err != nil {
		log.Println("Failed to stringify traits")
		return nil, err
	}
	if profile, err = toJsonString(profileRequest.Profile); err != nil {
		log.Println("Failed to stringify profile")
		return nil, err
	}
	return &db.UserProfileRequest{
		Traits:  traits,
		Profile: profile,
	}, nil
}

func toProfileResponse(profile *db.UserProfile) UserProfile {
	return UserProfile{
		Id:      profile.ID,
		Traits:  profile.Traits,
		Profile: profile.Profile,
	}
}
