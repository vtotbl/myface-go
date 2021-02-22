package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func getCurrentUserId(c *gin.Context) (int, error) {
	userId := c.MustGet("user_id")
	if nil == userId {
		return 0, errors.New("no current user")
	}
	id, _ := strconv.Atoi(userId.(string))

	return id, nil
}
