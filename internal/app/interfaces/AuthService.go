package interfaces

import (
	"github.com/Valeriy-Totubalin/myface-go/internal/delivery/http/request"
	"github.com/gin-gonic/gin"
)

type AuthService interface {
	SignUp(c *gin.Context, data request.SignUp) error
	SignIn(c *gin.Context, data request.SignIn) error
	Refresh(c *gin.Context, data request.Refresh) error
	LogOut(userId int) error
	IsExistsActiveSession(userId int) bool
}
