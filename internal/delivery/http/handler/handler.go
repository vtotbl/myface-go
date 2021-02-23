package handler

import (
	_ "github.com/Valeriy-Totubalin/myface-go/docs"
	"github.com/Valeriy-Totubalin/myface-go/pkg/token_manager"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"
)

type Handler struct {
	TokenManager token_manager.TokenManager
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		v1 := auth.Group("/v1")
		{
			v1.POST("/sign-up", h.signUp)
			v1.POST("/sign-in", h.signIn)
			v1.POST("/refresh", h.refresh)
			v1.POST("/log-out", h.logOut)
		}
	}

	api := router.Group("/api")
	{
		v1 := api.Group("/v1", h.checkToken)
		//v1.Use(h.checkToken) //middleware
		{
			v1.POST("/ping", h.pong)
			photo := v1.Group("/photo")
			{
				photo.POST("", h.upload)
				photo.GET("/:id", h.get)
			}
			rating := v1.Group("/rating")
			{
				rating.POST("", h.setRating)
			}
		}
	}

	return router
}
