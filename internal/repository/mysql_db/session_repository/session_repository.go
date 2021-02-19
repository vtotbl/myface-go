package session_repository

import (
	"errors"

	"github.com/Valeriy-Totubalin/myface-go/internal/domain"
	"github.com/Valeriy-Totubalin/myface-go/internal/repository/mysql_db"
	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	Id           int
	RefreshToken string
	ExpiresAt    string
	UserId       int
}

func CreateSession(session domain.Session) error {
	db, err := mysql_db.GetDB()
	if nil != err {
		return err
	}

	oldSession := Session{}
	db.Where("user_id = ?", session.UserId).Find(&oldSession)
	if 0 == oldSession.Id {
		db.Create(&Session{
			Id:           session.Id,
			RefreshToken: session.RefreshToken,
			ExpiresAt:    session.ExpiresAt,
			UserId:       session.UserId,
		})

		return nil
	}
	oldSession.RefreshToken = session.RefreshToken
	oldSession.ExpiresAt = session.ExpiresAt
	db.Save(oldSession)

	return nil
}

func GetByUserId(userId int) (*domain.Session, error) {
	db, err := mysql_db.GetDB()
	if nil != err {
		return nil, err
	}

	session := Session{}
	db.Where("user_id = ?", userId).Find(&session)
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
