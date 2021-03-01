package interfaces

import "github.com/Valeriy-Totubalin/myface-go/internal/domain"

type SessionRepository interface {
	CreateSession(session domain.Session) error
	GetByRefresh(token string) (*domain.Session, error)
	GetByUserId(userId int) (*domain.Session, error)
	DeleteByUserId(userId int) error
}
