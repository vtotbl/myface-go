package rating_service

import (
	"github.com/Valeriy-Totubalin/myface-go/internal/domain"
	"github.com/Valeriy-Totubalin/myface-go/internal/repository/mysql_db/photo_repository"
	"github.com/Valeriy-Totubalin/myface-go/internal/repository/mysql_db/rating_repository"
)

type RatingService struct {
	Repository *rating_repository.RatingRepository
}

func NewRatingService() (*RatingService, error) {
	repo, err := rating_repository.NewRatingRepository()
	if nil != err {
		return nil, err
	}
	service := RatingService{
		Repository: repo,
	}

	return &service, nil
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
	photo, err := photo_repository.GetById(photoId)
	if nil != err {
		return false, err
	}
	if photo.UserId == userId {
		return false, nil
	}

	return true, nil
}
