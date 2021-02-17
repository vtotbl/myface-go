package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) pong(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"message": "pong",
	})
}
