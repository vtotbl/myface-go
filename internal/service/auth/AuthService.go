package auth

import (
	"errors"
	"strconv"
	"time"

	"github.com/Valeriy-Totubalin/myface-go/internal/delivery/http/request"
	"github.com/Valeriy-Totubalin/myface-go/internal/domain"
	"github.com/Valeriy-Totubalin/myface-go/internal/repository/mysql_db/session_repository"
	"github.com/Valeriy-Totubalin/myface-go/internal/repository/mysql_db/user_repository"
	"github.com/Valeriy-Totubalin/myface-go/pkg/token_manager"
	"github.com/gin-gonic/gin"
)

type AuthService struct {
	PasswordHasher    *PasswordHasher
	UserRepository    *user_repository.UserRepository
	SessionRepository *session_repository.SessionRepository
}

func NewAuthService() (*AuthService, error) {
	passwordHasher, err := NewPasswordHasher()
	if nil != err {
		return nil, err
	}
	userRepo, err := user_repository.NewUserRepository()
	if nil != err {
		return nil, err
	}
	sessionRepo, err := session_repository.NewSessionRepository()
	if nil != err {
		return nil, err
	}

	service := AuthService{
		PasswordHasher:    passwordHasher,
		UserRepository:    userRepo,
		SessionRepository: sessionRepo,
	}

	return &service, nil
}

func (service *AuthService) SignUp(c *gin.Context, data request.SignUp) error {
	data.Password = service.PasswordHasher.GenerateHash(data.Password)
	user := domain.User{
		Login:    data.Login,
		Password: data.Password,
		Sex:      data.Sex,
	}

	err := service.UserRepository.SignUp(user)
	if nil != err {
		return err
	}

	return nil
}

func (service *AuthService) SignIn(c *gin.Context, data request.SignIn) error {
	user, err := service.UserRepository.GetByLogin(data.Login)
	if nil != err {
		return err
	}

	err = service.PasswordHasher.CheckPassword(data.Password, user.Password)
	if nil != err {
		return err
	}

	secret := c.MustGet("secret_key").(string)
	tokens, err := service.createSession(user.Id, secret)
	if nil != err {
		return err
	}

	c.Set("access_token", tokens.AccessToken)
	c.Set("refresh_token", tokens.RefreshToken)

	return nil
}

func (service *AuthService) Refresh(c *gin.Context, data request.Refresh) error {
	session, err := service.SessionRepository.GetByRefresh(data.Token)
	if nil != err {
		return err
	}

	if session.ExpiresAt < time.Now().String() {
		return errors.New("Token expired")
	}

	secret := c.MustGet("secret_key").(string)
	tokens, err := service.createSession(session.UserId, secret)
	if nil != err {
		return err
	}

	c.Set("access_token", tokens.AccessToken)
	c.Set("refresh_token", tokens.RefreshToken)

	return nil
}

func (service *AuthService) LogOut(userId int) error {
	err := service.SessionRepository.DeleteByUserId(userId)
	if nil != err {
		return err
	}

	return nil
}

func (service *AuthService) IsExistsActiveSession(userId int) bool {
	_, err := service.SessionRepository.GetByUserId(userId)
	if nil != err {
		return false
	}

	return true
}

func (service *AuthService) createSession(userId int, secret string) (token_manager.Tokens, error) {
	var res token_manager.Tokens

	tokenManager, err := token_manager.NewManager(secret)
	if nil != err {
		return res, err
	}

	res.AccessToken, err = tokenManager.NewJWT(strconv.Itoa(userId), 15*time.Minute) // 15 минут
	if nil != err {
		return res, err
	}

	res.RefreshToken, err = tokenManager.NewRefreshToken()
	if nil != err {
		return res, err
	}

	session := domain.Session{
		RefreshToken: res.RefreshToken,
		ExpiresAt:    time.Now().Add(720 * time.Hour).String(), // 30 дней
		UserId:       userId,
	}
	err = service.SessionRepository.CreateSession(session)
	if nil != err {
		return res, err
	}

	return res, nil
}
