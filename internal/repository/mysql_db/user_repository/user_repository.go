package user_repository

import (
	"errors"
	"log"

	"github.com/Valeriy-Totubalin/myface-go/internal/domain"
	"github.com/Valeriy-Totubalin/myface-go/internal/repository/mysql_db"
)

func SignUp(user domain.User) error {
	if isExists(user) {
		return errors.New("User already exists")
	}
	db := mysql_db.GetDB()
	db.Query("INSERT INTO `users` (`login`, `password`, `sex`) VALUES" + "(\"" + user.Login + "\", \"" + user.Password + "\", \"" + user.Sex + "\")")

	return nil
}

func isExists(user domain.User) bool {
	db := mysql_db.GetDB()
	rows := db.Query("SELECT `login` FROM `users` WHERE `login` = \"" + user.Login + "\"")
	if rows.Next() {
		return true
	}

	return false
}

func GetByLogin(login string) (domain.User, error) {
	db := mysql_db.GetDB()
	rows := db.Query("SELECT `id`, `login`, `password`, `sex` FROM `users` WHERE `login` = \"" + login + "\"")

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Login, &user.Password, &user.Sex)
		if err != nil {
			log.Fatal(err.Error())
		}

		return user, nil
	}

	return user, errors.New("User does not exist")
}
