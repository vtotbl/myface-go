package auth

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func generateHash(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
	}

	// err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password+"15"))
	// fmt.Println(err) // nil means it is a match

	return string(hashedPassword)
}
