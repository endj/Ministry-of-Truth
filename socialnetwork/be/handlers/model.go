package handlers

type UserProfileRequest struct {
	Traits  map[string]any `json:"traits"`
	Profile map[string]any `json:"profile"`
}

type UserProfile struct {
	Id      int    `json:"id"`
	Traits  string `json:"traits"`
	Profile string `json:"profile"`
}

type Post struct {
	Id        int    `json:"id"`
	CreatedAt int64  `json:"createdAt"`
	AuthordId int    `json:"authorId"`
	ThreadId  string `json:"threadId"`
	Content   string `json:"content"`
}

type PostRequest struct {
	AuthordId int    `json:"authorId"`
	ThreadId  string `json:"threadId"`
	Content   string `json:"content"`
}
