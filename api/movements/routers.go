package movements

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"vaultfitness.io/api/common"
)

func MovementsRegister(router *gin.RouterGroup) {

	router.GET("/", MovementsList)
}

type Result struct {
	ID           int
	Name         string
	Abbreviation string
}

func MovementsList(c *gin.Context) {
	db := common.GetDB()

	var result []Result
	db.Raw("SELECT id, name, abbreviation FROM movement_types").Scan(&result)

	c.JSON(http.StatusOK, result)
}
