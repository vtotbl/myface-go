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

func (repo *PhotoRepository) CreatePhoto(photo domain.Photo) (*domain.Photo, error) {
	db, err := mysql_db.GetDB()
	if nil != err {
		return nil, err
	}

	newPhoto := Photo{
		Path:   photo.Path,
		UserId: photo.UserId,
	}

	db.Create(&newPhoto)

	return &domain.Photo{
		Id:     newPhoto.Id,
		Path:   newPhoto.Path,
		UserId: newPhoto.UserId,
	}, nil
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

func (repo *PhotoRepository) GetRandom(userId int) (*domain.Photo, error) {
	db, err := mysql_db.GetDB()
	if nil != err {
		return nil, err
	}

	photo := Photo{}

	db.Order("RAND()").Limit(1).Model(&photo).Select("photos.id, photos.path, photos.user_id").Joins("LEFT JOIN ratings on ratings.photo_id = photos.id").Where("(ratings.user_id != ? OR ratings.user_id IS NULL) AND photos.user_id != ?", userId, userId).Scan(&photo)

	return &domain.Photo{
		Id:     photo.Id,
		Path:   photo.Path,
		UserId: photo.UserId,
	}, nil
}

func (repo *PhotoRepository) Delete(photoId int) error {
	db, err := mysql_db.GetDB()
	if nil != err {
		return err
	}
	photo := Photo{
		Id: photoId,
	}
	db.Delete(&photo)

	return nil
}
