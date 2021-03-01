package interfaces

import "github.com/Valeriy-Totubalin/myface-go/internal/domain"

type UserRepository interface {
	SignUp(user domain.User) error
	GetByLogin(login string) (*domain.User, error)
}
