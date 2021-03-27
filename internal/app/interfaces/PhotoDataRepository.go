package interfaces

import "github.com/Valeriy-Totubalin/myface-go/internal/domain"

type PhotoDataRepository interface {
	CreatePhoto(photo domain.Photo) (*domain.Photo, error)
	GetById(id int) (*domain.Photo, error)
	GetByUserId(userId int) ([]*domain.Photo, error)
	GetRandom(userId int) (*domain.Photo, error)
	Delete(photoId int) error
}
