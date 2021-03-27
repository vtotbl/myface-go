package handler

import (
	_ "github.com/Valeriy-Totubalin/myface-go/docs"
	"github.com/Valeriy-Totubalin/myface-go/internal/app/interfaces"
	"github.com/Valeriy-Totubalin/myface-go/pkg/token_manager"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"
)

const UnknowError = "Unknown error"
const PhotoNotFound = "Photo not found for this user"
const PhotoDeleted = "Photo deleted successfully"

type Handler struct {
	TokenManager   token_manager.TokenManager
	ServiceFactory interfaces.ServiceFactory
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
			v1.POST("/log-out", h.checkToken, h.logOut)
		}
	}

	api := router.Group("/api")
	{
		v1 := api.Group("/v1", h.checkToken)
		{
			v1.POST("/ping", h.pong)
			photo := v1.Group("/photo")
			{
				photo.POST("", h.upload)
				photo.PUT("", h.change)
				photo.GET("/random", h.getRandom)
				photo.GET("", h.get)
				photo.DELETE("", h.deletePhoto)
			}
			rating := v1.Group("/rating")
			{
				rating.POST("", h.setRating)
			}
		}
	}

	return router
}
