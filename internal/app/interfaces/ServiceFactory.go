package interfaces

type ServiceFactory interface {
	CreateAuthService() (AuthService, error)
	CreatePhotoService() (PhotoService, error)
	CreateRatingService() (RatingService, error)
}
