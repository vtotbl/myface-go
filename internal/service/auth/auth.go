package auth

import (
	"github.com/Valeriy-Totubalin/myface-go/internal/domain"
	"github.com/Valeriy-Totubalin/myface-go/internal/repository/mysql_db/user_repository"
)

func SignUp(user domain.User) error {
	user.Password = generateHash(user.Password)
	err := user_repository.SignUp(user)
	if nil != err {
		return err
	}

	return nil
}
