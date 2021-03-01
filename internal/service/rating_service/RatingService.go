package rating_service

import (
	"github.com/Valeriy-Totubalin/myface-go/internal/app/interfaces"
	"github.com/Valeriy-Totubalin/myface-go/internal/domain"
)

type RatingService struct {
	Repository      interfaces.RatingRepository
	PhotoRepository interfaces.PhotoDataRepository
}

func NewRatingService(
	repo interfaces.RatingRepository,
	photoRepo interfaces.PhotoDataRepository,
) interfaces.RatingService {

	return &RatingService{
		Repository:      repo,
		PhotoRepository: photoRepo,
	}
}

func (service *RatingService) SetRatingForPhoto(rating float64, photoId int, userId int) error {
	data := domain.Rating{
		Score:   rating,
		PhotoId: photoId,
		UserId:  userId,
	}
	err := service.Repository.CreateRating(&data)
	if nil != err {
		return err
	}

	return nil
}

func (service *RatingService) CanSetRatingForPhoto(userId int, photoId int) (bool, error) {
	photo, err := service.PhotoRepository.GetById(photoId)
	if nil != err {
		return false, err
	}
	if photo.UserId == userId {
		return false, nil
	}

	return true, nil
}
