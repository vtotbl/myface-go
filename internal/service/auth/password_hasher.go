package auth

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func generateHash(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
	}

	return string(hashedPassword)
}

func checkPassword(password string, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if nil != err {
		return errors.New("Incorrect login or password")
	}

	return nil
}
