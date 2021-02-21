package photo_repository

import (
	"errors"

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

func GetById(id int) (*domain.Photo, error) {
	db, err := mysql_db.GetDB()
	if nil != err {
		return nil, err
	}

	photo := Photo{
		Id: id,
	}

	db.Limit(1).Find(&photo)
	if 0 == photo.UserId {
		return nil, errors.New("Photo does not exist")
	}

	return &domain.Photo{
		Id:     photo.Id,
		Path:   photo.Path,
		UserId: photo.UserId,
	}, nil
}
