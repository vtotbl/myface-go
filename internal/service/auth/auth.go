package auth

import (
	"github.com/Valeriy-Totubalin/myface-go/internal/delivery/http/request"
	"github.com/Valeriy-Totubalin/myface-go/internal/domain"
	"github.com/Valeriy-Totubalin/myface-go/internal/repository/mysql_db/user_repository"
)

func SignUp(data request.SignUp) error {
	data.Password = generateHash(data.Password)
	user := domain.User{
		data.Login,
		data.Password,
		data.Sex,
	}

	err := user_repository.SignUp(user)
	if nil != err {
		return err
	}

	return nil
}

func SignIn(data request.SignIn) error {
	user, err := user_repository.GetByLogin(data.Login)
	if nil != err {
		return err
	}

	err = checkPassword(data.Password, user.Password)
	if nil != err {
		return err
	}

	//toDo нужно генерить jwt токен ну или с контроллера.

	return nil
}
