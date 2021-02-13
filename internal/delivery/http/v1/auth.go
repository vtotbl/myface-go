package v1

import (
	"github.com/Valeriy-Totubalin/myface-go/internal/delivery/http"
	"github.com/gin-gonic/gin"
)

func (h *http.Handler) signUp(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "test ok",
	})
}

func (h *http.Handler) signIn(c *gin.Context) {

}
