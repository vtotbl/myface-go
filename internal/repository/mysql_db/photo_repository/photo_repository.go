package photo_repository

import (
	"github.com/Valeriy-Totubalin/myface-go/internal/domain"
	"github.com/Valeriy-Totubalin/myface-go/internal/repository/mysql_db"
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	Id     int
	Path   string
	UserId int
}

func CreatePhoto(photo domain.Photo) error {
	db, err := mysql_db.GetDB()
	if nil != err {
		return err
	}

	db.Create(&Photo{
		Path:   photo.Path,
		UserId: photo.UserId,
	})

	return nil
}
