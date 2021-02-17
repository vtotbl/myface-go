package auth

import (
	"strconv"
	"time"

	"github.com/Valeriy-Totubalin/myface-go/internal/delivery/http/request"
	"github.com/Valeriy-Totubalin/myface-go/internal/domain"
	"github.com/Valeriy-Totubalin/myface-go/internal/repository/mysql_db/user_repository"
	"github.com/Valeriy-Totubalin/myface-go/pkg/token_manager"
	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context, data request.SignUp) error {
	data.Password = generateHash(data.Password)
	user := domain.User{
		Login:    data.Login,
		Password: data.Password,
		Sex:      data.Sex,
	}

	err := user_repository.SignUp(user)
	if nil != err {
		return err
	}

	return nil
}

func SignIn(c *gin.Context, data request.SignIn) error {
	user, err := user_repository.GetByLogin(data.Login)
	if nil != err {
		return err
	}

	err = checkPassword(data.Password, user.Password)
	if nil != err {
		return err
	}

	secret := c.MustGet("secret_key").(string)
	tokens, err := createSession(user.Id, secret)
	if nil != err {
		return err
	}

	c.Set("access_token", tokens.AccessToken)
	c.Set("refresh_token", tokens.RefreshToken)

	return nil
}

func createSession(userId int, secret string) (token_manager.Tokens, error) {
	var res token_manager.Tokens

	tokenManager, err := token_manager.NewManager(secret)
	if nil != err {
		return res, err
	}

	res.AccessToken, err = tokenManager.NewJWT(strconv.Itoa(userId), 15*time.Minute)
	if nil != err {
		return res, err
	}

	res.RefreshToken, err = tokenManager.NewRefreshToken()
	if nil != err {
		return res, err
	}

	return res, nil
}

// func CheckToken(accessToken string) (bool, error) {
// 	secret := c.MustGet("secret_key")

// 	tokenManager, err := token_manager.NewManager(secret)
// 	if nil != err {
// 		return false, err
// 	}

// 	test, err := tokenManager.Parse(accessToken)
// 	if nil != err {
// 		return false, err
// 	}

// 	fmt.Println(test)
// 	return true, nil
// }

// func getSecretKey() (string, error) {
// 	pwd, err := os.Getwd()
// 	if err != nil {
// 		return "", err
// 	}

// 	config, err := config_manager.GetJWTConfig(pwd + "/internal/config/jwt.json")
// 	if nil != err {
// 		return "", err
// 	}

// 	return config.SecretKey, nil
// }
