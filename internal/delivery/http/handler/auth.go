package handler

import (
	"net/http"

	"github.com/Valeriy-Totubalin/myface-go/internal/domain"
	"github.com/Valeriy-Totubalin/myface-go/internal/service/auth"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := auth.SignUp(user)
	if nil != err {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
	// c.JSON(200, gin.H{
	// 	"login":    user.Login,
	// 	"password": user.Password,
	// 	"sex":      user.Sex,
	// })
	// var json User
	// if err := c.ShouldBindJSON(&json); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// if json.Login != "manu" || json.Password != "123" {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}

func (h *Handler) signIn(c *gin.Context) {

}
