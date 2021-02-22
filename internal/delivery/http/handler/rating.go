package handler

import (
	"log"
	"net/http"

	"github.com/Valeriy-Totubalin/myface-go/internal/delivery/http/request"
	"github.com/Valeriy-Totubalin/myface-go/internal/service/rating_service"
	"github.com/gin-gonic/gin"
)

const UNKNOW_ERROR = "Unknown error"

func (h *Handler) setRating(c *gin.Context) {
	var data request.SetRatingInput
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	service, err := rating_service.NewRatingService()
	if nil != err {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": UNKNOW_ERROR})
		return
	}

	userId, _ := getCurrentUserId(c)

	canSet, err := service.CanSetRatingForPhoto(userId, data.PhotoId)
	if nil != err {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": UNKNOW_ERROR})
		return
	}

	if false == canSet {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You cannot change the rating of your photo"})
		return
	}

	err = service.SetRatingForPhoto(data.Rating, data.PhotoId, userId)
	if nil != err {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": UNKNOW_ERROR})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
