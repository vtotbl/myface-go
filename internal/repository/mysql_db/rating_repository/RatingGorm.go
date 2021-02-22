package rating_repository

import "gorm.io/gorm"

type Rating struct {
	gorm.Model
	Id      int
	Score   float64
	PhotoId int
	UserId  int
}
