package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Valeriy-Totubalin/myface-go/internal/delivery/http/request"
	"github.com/Valeriy-Totubalin/myface-go/internal/delivery/http/response"
	"github.com/Valeriy-Totubalin/myface-go/internal/service/photo_service"
	"github.com/gin-gonic/gin"
)

// @Summary upload
// @Security ApiKeyAuth
// @Tags api
// @Description upload photo to server
// @ID upload
// @Accept  json
// @Produce  json
// @Param input body request.Upload true "Base64 encoded photo"
// @Success 200 {object} response.Message
// @Failure 400 {object} response.Error
// @Failure 500 {object} response.Error
// @Router /api/v1/photo [post]
func (h *Handler) upload(c *gin.Context) {
	var data request.Upload
	if err := c.ShouldBindJSON(&data); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Error{Error: err.Error()})
		return
	}

	userId := c.MustGet("user_id")
	if nil == userId {
		return
	}
	id, _ := strconv.Atoi(userId.(string))

	service, err := photo_service.NewPhotoService()
	if nil != err {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": UnknowError})
		return
	}

	err = service.CheckCorrectData(data.Photo)
	if nil != err {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Error{Error: err.Error()})
		return
	}
	err = service.Upload(id, data.Photo)
	if nil != err {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Error{Error: UnknowError})
		return
	}

	c.JSON(http.StatusOK, response.Message{Message: "ok"})
}

// @Summary get photo
// @Security ApiKeyAuth
// @Tags api
// @Description get photo by id
// @ID get-photo
// @Accept  json
// @Produce  json
// @Param id path int true "Photo ID"
// @Success 200 {object} response.GetPhoto
// @Failure 400 {object} response.Error
// @Failure 401 {object} response.Error
// @Failure 500 {object} response.Error
// @Router /api/v1/photo/{id} [get]
func (h *Handler) get(c *gin.Context) {
	var photoInput request.Photo
	if err := c.ShouldBindUri(&photoInput); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Error{Error: err.Error()})
		return
	}

	id, err := strconv.Atoi(photoInput.Id)
	if nil != err {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Error{Error: UnknowError})
		return
	}

	userId := c.MustGet("user_id")
	if nil == userId {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Error{Error: "Invalid user id"})
		return
	}
	userIdInt, _ := strconv.Atoi(userId.(string))

	service, err := photo_service.NewPhotoService()
	if nil != err {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Error{Error: UnknowError})
		return
	}

	canGet, err := service.CanGet(userIdInt, id)
	if nil != err {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Error{Error: UnknowError})
		return
	}

	if false == canGet {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Error{Error: "Photo not found for this user"})
		return
	}

	base64, err := service.GetById(id)
	if nil != err {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Error{Error: UnknowError})
		return
	}

	c.JSON(http.StatusOK, response.GetPhoto{
		PhotoId: photoInput.Id,
		Base64:  base64,
	})
}

// @Summary get photos
// @Security ApiKeyAuth
// @Tags api
// @Description get all photos of the current user
// @ID get-photos
// @Accept  json
// @Produce  json
// @Success 200 {object} []response.GetPhoto
// @Failure 401 {object} response.Error
// @Failure 500 {object} response.Error
// @Router /api/v1/photo [get]
func (h *Handler) getAll(c *gin.Context) {
	userId, err := getCurrentUserId(c)
	if nil != err {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Error{Error: UnknowError})
		return
	}

	service, err := photo_service.NewPhotoService()
	if nil != err {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Error{Error: UnknowError})
		return
	}

	photos, err := service.GetByUserId(userId)
	if nil != err {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Error{Error: UnknowError})
		return
	}

	var base64Photos []*response.GetPhoto
	for _, photo := range photos {
		base64Photos = append(base64Photos, &response.GetPhoto{
			PhotoId: strconv.Itoa(photo.Id),
			Base64:  photo.Base64,
		})
	}
	c.JSON(http.StatusOK, base64Photos)
}
