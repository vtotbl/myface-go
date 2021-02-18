package handler

import (
	"net/http"

	"github.com/Valeriy-Totubalin/myface-go/internal/delivery/http/request"
	"github.com/Valeriy-Totubalin/myface-go/internal/service/auth"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var data request.SignUp
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Set("secret_key", h.TokenManager.GetSecretKey())
	err := auth.SignUp(c, data)
	if nil != err {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	signInInput := request.SignIn{
		data.Login,
		data.Password,
	}
	err = auth.SignIn(c, signInInput)
	if nil != err {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"access_token":  c.MustGet("access_token"),
			"refresh_token": c.MustGet("refresh_token"),
		})
	}
}

func (h *Handler) signIn(c *gin.Context) {
	var data request.SignIn
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Set("secret_key", h.TokenManager.GetSecretKey())
	err := auth.SignIn(c, data)
	if nil != err {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"access_token":  c.MustGet("access_token"),
			"refresh_token": c.MustGet("refresh_token"),
		})
	}
}
