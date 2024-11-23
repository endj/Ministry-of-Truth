package handlers

import (
	"app/db"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAdminGetUserProfileHandler(t *testing.T) {
	repo := &db.MockUserRepository{}
	handler := AdmingListUsersHandler(repo)

	// Create the request
	req, err := http.NewRequest("GET", "/profiles", nil)
	if err != nil {
		t.Fatal("Failed to create request:", err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	var response []UserFullProfileResponse
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatal("Failed to decode response:", err)
	}

	if len(response) != 1 {
		t.Errorf("Expected 1 profile, got %d", len(response))
	}

	if response[0].Name != "John Doe" {
		t.Errorf("Expected name 'John Doe', got '%s'", response[0].Name)
	}

	if response[0].Info != "Software Engineer" {
		t.Errorf("Expected info 'Software Engineer', got '%s'", response[0].Info)
	}

	if response[0].Age != 30 {
		t.Errorf("Expected age 30, got %d", response[0].Age)
	}
}

func TestAdminPostProfileHandler(t *testing.T) {
	repo := &db.MockUserRepository{}
	handler := AdminCreateProfileHandler(repo)

	// Create a test profile request
	profileRequest := UserFullProfileRequest{
		ExternalID:            12345,
		Name:                  "John Doe",
		Info:                  "Software Engineer",
		Age:                   30,
		Occupation:            "Engineer",
		PersonalityType:       "Introvert",
		PersonalityHumor:      "Dry",
		PersonalitySocial:     "Reserved",
		InterestsPrimary:      []string{"Coding", "Music"},
		InterestsSecondary:    []string{"Hiking"},
		InterestsPreferred:    []string{"Tech", "Science"},
		InterestsDisliked:     []string{"Politics"},
		BackgroundHometown:    "New York",
		BackgroundEducation:   "Bachelors",
		BackgroundValues:      []string{"Integrity", "Innovation"},
		CommStyleFavorites:    []string{"Efficiency", "Clarity"},
		CommStyleFormality:    "Formal",
		CommStyleConversation: "Structured",
		SocialGroups:          []string{"Work", "Family"},
		SocialFriendliness:    "High",
	}

	reqBody, err := json.Marshal(profileRequest)
	if err != nil {
		t.Fatal("Failed to marshal request body:", err)
	}

	req, err := http.NewRequest("POST", "/profile", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal("Failed to create request:", err)
	}

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, rr.Code)
	}

	var response UserFullProfileResponse
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatal("Failed to decode response:", err)
	}

	if response.Name != "John Doe" {
		t.Errorf("Expected name 'John Doe', got '%s'", response.Name)
	}

	if response.Info != "Software Engineer" {
		t.Errorf("Expected info 'Software Engineer', got '%s'", response.Info)
	}

	if response.Age != 30 {
		t.Errorf("Expected age 30, got %d", response.Age)
	}
}
