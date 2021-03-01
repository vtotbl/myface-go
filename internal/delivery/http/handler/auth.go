package handler

import (
	"net/http"
	"strconv"

	"github.com/Valeriy-Totubalin/myface-go/internal/delivery/http/request"
	"github.com/Valeriy-Totubalin/myface-go/internal/delivery/http/response"
	"github.com/gin-gonic/gin"
)

// @Summary sign-up
// @Tags auth
// @Description Create new account
// @ID sign-up
// @Accept  json
// @Produce  json
// @Param input body request.SignUp true "account info"
// @Success 201 {object} response.Tokens
// @Failure 400 {object} response.Error
// @Failure 500 {object} response.Error
// @Router /auth/v1/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var data request.SignUp
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, response.Error{Error: err.Error()})
		return
	}
	c.Set("secret_key", h.TokenManager.GetSecretKey())
	service, err := h.ServiceFactory.CreateAuthService()
	if nil != err {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Error{
			Error: UnknowError,
		})
		return
	}

	err = service.SignUp(c, data)
	if nil != err {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Error{Error: err.Error()})
		return
	}

	signInInput := request.SignIn{
		data.Login,
		data.Password,
	}
	err = service.SignIn(c, signInInput)
	if nil != err {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response.Tokens{
		AccessToken:  c.MustGet("access_token").(string),
		RefreshToken: c.MustGet("refresh_token").(string),
	})
}

// @Summary sign-in
// @Tags auth
// @Description Log in with an existing account
// @ID sign-in
// @Accept  json
// @Produce  json
// @Param input body request.SignIn true "login and password from the account"
// @Success 200 {object} response.Tokens
// @Failure 400 {object} response.Error
// @Failure 500 {object} response.Error
// @Router /auth/v1/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var data request.SignIn
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, response.Error{Error: err.Error()})
		return
	}
	c.Set("secret_key", h.TokenManager.GetSecretKey())

	service, err := h.ServiceFactory.CreateAuthService()
	if nil != err {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Error{
			Error: UnknowError,
		})
		return
	}

	err = service.SignIn(c, data)
	if nil != err {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  c.MustGet("access_token"),
		"refresh_token": c.MustGet("refresh_token"),
	})
}

// @Summary refresh
// @Tags auth
// @Description Refresh tokens
// @ID refresh
// @Accept  json
// @Produce  json
// @Param input body request.Refresh true "refresh token"
// @Success 200 {object} response.Tokens
// @Failure 400 {object} response.Error
// @Failure 500 {object} response.Error
// @Router /auth/v1/refresh [post]
func (h *Handler) refresh(c *gin.Context) {
	var data request.Refresh
	if err := c.ShouldBindJSON(&data); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Error{Error: err.Error()})
		return
	}

	c.Set("secret_key", h.TokenManager.GetSecretKey())

	service, err := h.ServiceFactory.CreateAuthService()
	if nil != err {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Error{
			Error: UnknowError,
		})
		return
	}

	err = service.Refresh(c, data)
	if nil != err {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Error{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  c.MustGet("access_token"),
		"refresh_token": c.MustGet("refresh_token"),
	})
}

// @Summary log-out
// @Security ApiKeyAuth
// @Tags auth
// @Description Log out
// @ID log-out
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Message
// @Failure 400 {object} response.Error
// @Failure 500 {object} response.Error
// @Router /auth/v1/log-out [post]
func (h *Handler) logOut(c *gin.Context) {
	h.checkToken(c)
	userId := c.MustGet("user_id")
	if nil == userId {
		return
	}
	id, _ := strconv.Atoi(userId.(string))

	service, err := h.ServiceFactory.CreateAuthService()
	if nil != err {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Error{
			Error: UnknowError,
		})
		return
	}

	if false == service.IsExistsActiveSession(id) {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Error{
			Error: "No active session",
		})
		return
	}

	err = service.LogOut(id)

	if nil != err {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Error{
			Error: UnknowError,
		})
		return
	}

	c.JSON(http.StatusOK, response.Message{Message: "ok"})
}
