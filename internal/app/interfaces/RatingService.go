package interfaces

type RatingService interface {
	SetRatingForPhoto(rating float64, photoId int, userId int) error
	CanSetRatingForPhoto(userId int, photoId int) (bool, error)
}
