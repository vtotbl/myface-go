package handler

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		v1 := auth.Group("/v1")
		{
			v1.POST("/sign-up", h.signUp)
			v1.POST("/sign-in", h.signIn)
		}
	}

	// api := router.Group("/api")
	// {

	// }

	return router
}
