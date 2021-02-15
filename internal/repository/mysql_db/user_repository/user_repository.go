package user_repository

import (
	"errors"

	"github.com/Valeriy-Totubalin/myface-go/internal/domain"
	"github.com/Valeriy-Totubalin/myface-go/internal/repository/mysql_db"
)

func SignUp(user domain.User) error {
	if isExists(user) {
		return errors.New("User already exists")
	}
	db := mysql_db.GetDB()
	db.Query("INSERT INTO `users` VALUES" + "(\"" + user.Login + "\", \"" + user.Password + "\", \"" + user.Sex + "\")")

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
