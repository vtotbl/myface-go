package photo_repository

import "gorm.io/gorm"

type Photo struct {
	gorm.Model
	Id     int
	Path   string
	UserId int
}
