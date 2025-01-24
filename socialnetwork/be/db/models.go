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
