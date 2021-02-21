package handler

import (
	"net/http"
	"strconv"

	"github.com/Valeriy-Totubalin/myface-go/internal/delivery/http/request"
	"github.com/Valeriy-Totubalin/myface-go/internal/service/photo"
	"github.com/gin-gonic/gin"
)

func (h *Handler) upload(c *gin.Context) {
	var data request.Upload
	if err := c.ShouldBindJSON(&data); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userId := c.MustGet("user_id")
	if nil == userId {
		return
	}
	id, _ := strconv.Atoi(userId.(string))

	err := photo.CheckCorrectData(data.Photo)

	if nil != err {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = photo.Upload(id, data.Photo)
	if nil != err {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
