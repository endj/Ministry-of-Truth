package db

type UserProfile struct {
	ID      int
	Name    string
	Traits  string
	Profile string
}

type UserProfileRequest struct {
	Name    string
	Traits  string
	Profile string
}

type Post struct {
	ID        int
	CreatedAt int64
	ThreadId  string
	AuthordId int
	Author    string
	OP        int
	Content   string
}

type PostRequest struct {
	AuthordId int
	ThreadId  string
	Content   string
}
