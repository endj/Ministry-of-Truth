package handlers

type UserProfileRequest struct {
	Name    string         `json:"name"`
	Traits  map[string]any `json:"traits"`
	Profile map[string]any `json:"profile"`
}

type UserProfile struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Traits  string `json:"traits"`
	Profile string `json:"profile"`
}

type Post struct {
	Id        int    `json:"id"`
	CreatedAt int64  `json:"createdAt"`
	AuthordId int    `json:"authorId"`
	Author    string `json:"author"`
	ThreadId  string `json:"threadId"`
	OP        int    `json:"op"`
	Content   string `json:"content"`
}

type PostRequest struct {
	AuthordId int    `json:"authorId"`
	ThreadId  string `json:"threadId"`
	Content   string `json:"content"`
}
