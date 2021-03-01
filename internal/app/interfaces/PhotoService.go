package interfaces

import "github.com/Valeriy-Totubalin/myface-go/internal/domain"

type PhotoService interface {
	Upload(userId int, data string) error
	CheckCorrectData(data string) error
	GetById(id int) (string, error)
	CanGet(userId, photoId int) (bool, error)
	GetByUserId(userId int) ([]*domain.PhotoBase64, error)
	GetRandom(userId int) (*domain.PhotoBase64, error)
}
