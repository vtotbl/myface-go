package session_repository

import "gorm.io/gorm"

type Session struct {
	gorm.Model
	Id           int
	RefreshToken string
	ExpiresAt    string
	UserId       int
}
