package domain

type Session struct {
	Id           int
	RefreshToken string
	ExpiresAt    string
	UserId       int
}
