package handler

import (
	"log"
	"net/http"

	"github.com/Valeriy-Totubalin/myface-go/internal/delivery/http/request"
	"github.com/Valeriy-Totubalin/myface-go/internal/delivery/http/response"
	"github.com/gin-gonic/gin"
)

// @Summary set-rating
// @Security ApiKeyAuth
// @Tags api
// @Description Set rating for photo
// @ID set-rating
// @Accept  json
// @Produce  json
// @Param input body request.SetRatingInput true "Rating from 1 to 10 and photo ID"
// @Success 200 {object} response.Message
// @Failure 400 {object} response.Error
// @Failure 500 {object} response.Error
// @Router /api/v1/rating [post]
func (h *Handler) setRating(c *gin.Context) {
	var data request.SetRatingInput
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, response.Error{Error: err.Error()})
		return
	}
	service, err := h.ServiceFactory.CreateRatingService()
	if nil != err {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Error{Error: UnknowError})
		return
	}

	userId, _ := getCurrentUserId(c)

	canSet, err := service.CanSetRatingForPhoto(userId, data.PhotoId)
	if nil != err {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Error{Error: UnknowError})
		return
	}

	if false == canSet {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Error{Error: "You cannot change the rating of your photo"})
		return
	}

	err = service.SetRatingForPhoto(data.Rating, data.PhotoId, userId)
	if nil != err {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Error{Error: UnknowError})
		return
	}

	c.JSON(http.StatusOK, response.Message{Message: "ok"})
}
