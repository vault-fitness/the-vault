package movements

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"vaultfitness.io/api/common"
)

func MovementsRegister(router *gin.RouterGroup) {
	router.GET("/type", MovementTypesList)
	router.POST("/type", CreateMovementType)
}

type createMovementTypeRequest struct {
	Name         string
	Abbreviation string
}

func CreateMovementType(c *gin.Context) {
	db := common.GetDB()

	var r createMovementTypeRequest

	var existingMovementType []MovementType
	db.Raw("SELECT id, name, abbreviation FROM movement_types WHERE name = $1 OR abbreviation = $2", r.Name, r.Abbreviation).Scan(&existingMovementType)

	if len(existingMovementType) > 0 {
		fmt.Println("Found movement already")
		c.String(http.StatusBadRequest, "Movement Type Already Exists")
		return
	}

	mt := MovementType{
		Name:         r.Name,
		Abbreviation: r.Abbreviation,
	}

	db.Create(&mt)

	fmt.Printf("Created movement type %d", mt.ID)

	c.JSON(http.StatusOK, mt)
}

func MovementTypesList(c *gin.Context) {
	db := common.GetDB()

	var mt []MovementType
	db.Raw("SELECT id, name, abbreviation FROM movement_types").Scan(&mt)

	c.JSON(http.StatusOK, mt)
}
