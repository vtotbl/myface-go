package user_repository

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       int
	Login    string
	Password string
	Sex      string
}
