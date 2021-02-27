package photo_repository

import (
	"errors"

	"github.com/Valeriy-Totubalin/myface-go/internal/domain"
	"github.com/Valeriy-Totubalin/myface-go/internal/repository/mysql_db"
)

type PhotoRepository struct {
}

func NewPhotoRepository() (*PhotoRepository, error) {
	return &PhotoRepository{}, nil
}

func (repo *PhotoRepository) CreatePhoto(photo domain.Photo) error {
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

func (repo *PhotoRepository) GetById(id int) (*domain.Photo, error) {
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

func (repo *PhotoRepository) GetByUserId(userId int) ([]*domain.Photo, error) {
	db, err := mysql_db.GetDB()
	if nil != err {
		return nil, err
	}

	photos := []Photo{}

	db.Where("user_id = ?", userId).Find(&photos)

	var domainPhotos []*domain.Photo
	for _, photo := range photos {
		domainPhotos = append(domainPhotos, &domain.Photo{
			Id:     photo.Id,
			Path:   photo.Path,
			UserId: photo.UserId,
		})
	}

	return domainPhotos, nil
}
