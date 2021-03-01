package password_hasher

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type PasswordHasher struct {
}

func NewPasswordHasher() (*PasswordHasher, error) {
	return &PasswordHasher{}, nil
}

func (service *PasswordHasher) GenerateHash(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
	}

	return string(hashedPassword)
}

func (service *PasswordHasher) CheckPassword(password string, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if nil != err {
		return errors.New("Incorrect login or password")
	}

	return nil
}
