package user_repository

import (
	"errors"

	"gorm.io/gorm"

	"github.com/Valeriy-Totubalin/myface-go/internal/domain"
	"github.com/Valeriy-Totubalin/myface-go/internal/repository/mysql_db"
)

type User struct {
	gorm.Model
	Id       int
	Login    string
	Password string
	Sex      string
}

func SignUp(user domain.User) error {
	db, err := mysql_db.GetDB()
	if nil != err {
		return err
	}

	if isExists(user) {
		return errors.New("User already exists")
	}
	db.Create(&User{
		Login:    user.Login,
		Password: user.Password,
		Sex:      user.Sex,
	})

	return nil
}

func isExists(user domain.User) bool {
	domainUser, _ := GetByLogin(user.Login)
	if 0 != domainUser.Id {
		return true
	}
	return false
}

func GetByLogin(login string) (*domain.User, error) {
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
