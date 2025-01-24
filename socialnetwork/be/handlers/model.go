package handlers

type UserProfileRequest struct {
	Traits  map[string]interface{} `json:"traits"`
	Profile map[string]interface{} `json:"profile"`
}

type UserProfile struct {
	Id      int    `json:"id"`
	Traits  string `json:"traits"`
	Profile string `json:"profile"`
}
