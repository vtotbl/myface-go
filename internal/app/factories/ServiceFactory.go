package factories

import (
	"github.com/Valeriy-Totubalin/myface-go/internal/app/interfaces"
	"github.com/Valeriy-Totubalin/myface-go/internal/repository/mysql_db/photo_repository"
	"github.com/Valeriy-Totubalin/myface-go/internal/repository/mysql_db/rating_repository"
	"github.com/Valeriy-Totubalin/myface-go/internal/repository/mysql_db/session_repository"
	"github.com/Valeriy-Totubalin/myface-go/internal/repository/mysql_db/user_repository"
	"github.com/Valeriy-Totubalin/myface-go/internal/repository/os/photo_os_repository"
	"github.com/Valeriy-Totubalin/myface-go/internal/service/auth"
	"github.com/Valeriy-Totubalin/myface-go/internal/service/photo_service"
	"github.com/Valeriy-Totubalin/myface-go/internal/service/rating_service"
	"github.com/Valeriy-Totubalin/myface-go/pkg/password_hasher"
)

type ServiceFactory struct {
}

func NewServiceFactory() (interfaces.ServiceFactory, error) {
	return &ServiceFactory{}, nil
}

func (factory *ServiceFactory) CreateAuthService() (interfaces.AuthService, error) {
	passwordHasher, err := password_hasher.NewPasswordHasher()
	if nil != err {
		return nil, err
	}
	userRepo, err := user_repository.NewUserRepository()
	if nil != err {
		return nil, err
	}
	sessionRepo, err := session_repository.NewSessionRepository()
	if nil != err {
		return nil, err
	}

	service := auth.NewAuthService(passwordHasher, userRepo, sessionRepo)

	return service, nil
}

func (factory *ServiceFactory) CreatePhotoService() (interfaces.PhotoService, error) {
	repo, err := photo_repository.NewPhotoRepository()
	if nil != err {
		return nil, err
	}

	repoFile, err := photo_os_repository.NewPhotoOsRepository()
	if nil != err {
		return nil, err
	}

	return photo_service.NewPhotoService(repo, repoFile), nil
}

func (factory *ServiceFactory) CreateRatingService() (interfaces.RatingService, error) {
	repo, err := rating_repository.NewRatingRepository()
	if nil != err {
		return nil, err
	}

	photoRepo, err := photo_repository.NewPhotoRepository()
	if nil != err {
		return nil, err
	}

	return rating_service.NewRatingService(repo, photoRepo), nil
}
