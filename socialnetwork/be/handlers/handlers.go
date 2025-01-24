package handlers

import (
	"app/db"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetProfilesHandler(w http.ResponseWriter, r *http.Request, repo db.UserRepository) {
	setHeaders(w)
	profiles, err := repo.QueryProfiles()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error querying database: %v", err), http.StatusInternalServerError)
		return
	}

	var responseProfiles []UserProfile = make([]UserProfile, 0)
	for _, profile := range profiles {
		responseProfiles = append(responseProfiles, toResponse(profile))
	}

	if err := json.NewEncoder(w).Encode(responseProfiles); err != nil {
		http.Error(w, fmt.Sprintf("Error encoding response: %v", err), http.StatusInternalServerError)
	}
}

func CreateProfileHandler(w http.ResponseWriter, r *http.Request, repo db.UserRepository) {
	setHeaders(w)
	var profileRequest UserProfileRequest

	if err := json.NewDecoder(r.Body).Decode(&profileRequest); err != nil {
		http.Error(w, "Invalid JSON payload: "+err.Error(), http.StatusBadRequest)
		return
	}
	if profileRequest.Profile == nil || profileRequest.Traits == nil {
		http.Error(w, "Invalid JSON payload: mising fields", http.StatusBadRequest)

		return
	}
	fmt.Println("Got request", profileRequest)

	traits, err := MapToString(profileRequest.Traits)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error processing request: %v", err), http.StatusInternalServerError)
		return
	}

	profile, err := MapToString(profileRequest.Profile)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error processing request: %v", err), http.StatusInternalServerError)
		return
	}
	profileResponse, err := repo.CreateProfile(db.UserProfileRequest{
		Traits:  traits,
		Profile: profile,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("Error processing request: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(toResponse(profileResponse)); err != nil {
		http.Error(w, fmt.Sprintf("Error encoding response: %v", err), http.StatusInternalServerError)
	}
}

func setHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Content-Type", "application/json")
}

func MapToString(m map[string]interface{}) (string, error) {
	bytes, err := json.Marshal(m)
	if err != nil {
		return "", fmt.Errorf("failed to marshal map to JSON: %w", err)
	}
	return string(bytes), nil
}

func StringToMap(str string) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(str), &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal string to map: %w", err)
	}
	return result, nil
}

func toResponse(profile db.UserProfile) UserProfile {
	return UserProfile{
		Id:      profile.ID,
		Traits:  profile.Traits,
		Profile: profile.Profile,
	}
}

func (u *UserProfileRequest) Validate() error {
	return nil
}
