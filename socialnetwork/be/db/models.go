package db

type UserProfile struct {
	ID      int
	Traits  string
	Profile string
}

type UserProfileRequest struct {
	Traits  string
	Profile string
}

type Post struct {
	ID        int
	CreatedAt int64
	ThreadId  string
	AuthordId int
	Content   string
}

type PostRequest struct {
	AuthordId int
	ThreadId  string
	Content   string
}
