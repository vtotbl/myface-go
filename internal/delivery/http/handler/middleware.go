package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const authorizationHeader = "Authorization"

func (h *Handler) checkToken(c *gin.Context) {
	id, err := h.parseAuthHeader(c)
	if nil != err {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
	}
	c.Set("user_id", id)
}

func (h *Handler) parseAuthHeader(c *gin.Context) (string, error) {
	header := c.GetHeader(authorizationHeader)
	if "" == header {
		return "", errors.New("Empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", errors.New("Invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return "", errors.New("Token is empty")
	}

	id, err := h.TokenManager.Parse(headerParts[1])
	if nil != err {
		return "", err
	}

	return id, nil
}
