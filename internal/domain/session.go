package domain

type Session struct {
	Id           int
	RefreshToken string
	ExpiresAt    int64
	UserId       int
}
