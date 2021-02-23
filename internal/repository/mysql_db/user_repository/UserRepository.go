package user_repository

import (
	"errors"

	"github.com/Valeriy-Totubalin/myface-go/internal/domain"
	"github.com/Valeriy-Totubalin/myface-go/internal/repository/mysql_db"
)

type UserRepository struct {
}

func NewUserRepository() (*UserRepository, error) {
	return &UserRepository{}, nil
}

func (repo *UserRepository) SignUp(user domain.User) error {
	db, err := mysql_db.GetDB()
	if nil != err {
		return err
	}

	if repo.isExists(user) {
		return errors.New("User already exists")
	}
	db.Create(&User{
		Login:    user.Login,
		Password: user.Password,
		Sex:      user.Sex,
	})

	return nil
}

func (repo *UserRepository) isExists(user domain.User) bool {
	domainUser, _ := repo.GetByLogin(user.Login)
	if 0 != domainUser.Id {
		return true
	}
	return false
}

func (repo *UserRepository) GetByLogin(login string) (*domain.User, error) {
	db, err := mysql_db.GetDB()
	if nil != err {
		return nil, err
	}
	user := User{}
	db.Where("login = ?", login).Find(&user)
	if 0 == user.Id {
		return &domain.User{}, errors.New("User does not exist")
	}
	return &domain.User{
		Id:       user.Id,
		Login:    user.Login,
		Password: user.Password,
		Sex:      user.Sex,
	}, nil
}
