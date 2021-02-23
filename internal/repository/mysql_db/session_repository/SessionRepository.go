package session_repository

import (
	"errors"

	"github.com/Valeriy-Totubalin/myface-go/internal/domain"
	"github.com/Valeriy-Totubalin/myface-go/internal/repository/mysql_db"
)

type SessionRepository struct {
}

func NewSessionRepository() (*SessionRepository, error) {
	return &SessionRepository{}, nil
}

func (repo *SessionRepository) CreateSession(session domain.Session) error {
	db, err := mysql_db.GetDB()
	if nil != err {
		return err
	}

	oldSession := Session{}
	db.Where("user_id = ?", session.UserId).Find(&oldSession)

	if 0 != oldSession.Id {
		db.Delete(&oldSession)
	}

	db.Create(&Session{
		Id:           session.Id,
		RefreshToken: session.RefreshToken,
		ExpiresAt:    session.ExpiresAt,
		UserId:       session.UserId,
	})

	return nil
}

func (repo *SessionRepository) GetByRefresh(token string) (*domain.Session, error) {
	db, err := mysql_db.GetDB()
	if nil != err {
		return nil, err
	}

	session := Session{}
	db.Where("refresh_token = ?", token).Find(&session)
	if 0 == session.Id {
		return nil, errors.New("Session does not exist")
	}

	return &domain.Session{
		Id:           session.Id,
		RefreshToken: session.RefreshToken,
		ExpiresAt:    session.ExpiresAt,
		UserId:       session.UserId,
	}, nil
}

func (repo *SessionRepository) DeleteByUserId(userId int) error {
	db, err := mysql_db.GetDB()
	if nil != err {
		return err
	}

	session := Session{}
	db.Where("user_id = ?", userId).Find(&session)

	if 0 == session.Id {
		return errors.New("Session does not exist")
	}
	db.Delete(&session)

	return nil
}
