package handlers

import (
	"app/db"
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleGetProfiles(w http.ResponseWriter, r *http.Request) {
	rows, err := db.QueryFullProfiles()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error querying database: %v", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var profiles []UserFullProfileResponse
	for rows.Next() {
		var profile UserFullProfileResponse
		var personality, interests, background, communicationStyle, socialConnections []byte

		err := rows.Scan(
			&profile.ID,
			&profile.ExternalID,
			&profile.Name,
			&profile.Info,
			&profile.Age,
			&profile.Occupation,
			&personality,
			&interests,
			&background,
			&communicationStyle,
			&socialConnections,
		)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error scanning row: %v", err), http.StatusInternalServerError)
			return
		}

		if err := json.Unmarshal(personality, &profile.Personality); err != nil {
			http.Error(w, fmt.Sprintf("Error unmarshaling personality: %v", err), http.StatusInternalServerError)
			return
		}
		if err := json.Unmarshal(interests, &profile.Interests); err != nil {
			http.Error(w, fmt.Sprintf("Error unmarshaling interests: %v", err), http.StatusInternalServerError)
			return
		}
		if err := json.Unmarshal(background, &profile.Background); err != nil {
			http.Error(w, fmt.Sprintf("Error unmarshaling background: %v", err), http.StatusInternalServerError)
			return
		}
		if err := json.Unmarshal(communicationStyle, &profile.CommunicationStyle); err != nil {
			http.Error(w, fmt.Sprintf("Error unmarshaling communication style: %v", err), http.StatusInternalServerError)
			return
		}
		if err := json.Unmarshal(socialConnections, &profile.SocialConnections); err != nil {
			http.Error(w, fmt.Sprintf("Error unmarshaling social connections: %v", err), http.StatusInternalServerError)
			return
		}

		profiles = append(profiles, profile)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, fmt.Sprintf("Error iterating over rows: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(profiles); err != nil {
		http.Error(w, fmt.Sprintf("Error encoding response: %v", err), http.StatusInternalServerError)
		return
	}
}

func handlePostProfile(w http.ResponseWriter, r *http.Request) {
	var profileRequest db.UserProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&profileRequest); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	profileResponse, err := db.InsertUserProfileWithTemplate(profileRequest)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error processing request: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profileResponse)
}
