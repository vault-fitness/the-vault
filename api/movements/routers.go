package movements

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MovementsRegister(router *gin.RouterGroup) {
	router.GET("/", MovementsList)
}

func MovementsList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"movements": nil})
}
