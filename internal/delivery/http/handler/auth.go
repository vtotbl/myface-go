package handler

import (
	"net/http"
	"strconv"

	"github.com/Valeriy-Totubalin/myface-go/internal/delivery/http/request"
	"github.com/Valeriy-Totubalin/myface-go/internal/service/auth"
	"github.com/gin-gonic/gin"
)

// ShowAccount godoc
// @Summary SignUp
// @Tags auth
// @Description create new account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body request.SignUp true "account info"
// @Success 200 {object} string
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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  c.MustGet("access_token"),
		"refresh_token": c.MustGet("refresh_token"),
	})
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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  c.MustGet("access_token"),
		"refresh_token": c.MustGet("refresh_token"),
	})
}

func (h *Handler) refresh(c *gin.Context) {
	var data request.Refresh
	if err := c.ShouldBindJSON(&data); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Set("secret_key", h.TokenManager.GetSecretKey())
	err := auth.Refresh(c, data)
	if nil != err {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  c.MustGet("access_token"),
		"refresh_token": c.MustGet("refresh_token"),
	})
}

func (h *Handler) logOut(c *gin.Context) {
	h.checkToken(c)
	userId := c.MustGet("user_id")
	if nil == userId {
		return
	}
	id, _ := strconv.Atoi(userId.(string))
	err := auth.LogOut(id)

	if nil != err {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
