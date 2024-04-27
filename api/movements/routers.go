package movements

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"vaultfitness.io/api/common"
)

func MovementsRegister(router *gin.RouterGroup) {
	router.GET("/type", MovementTypesList)
	router.POST("/type", CreateMovementType)
}

type MovementType struct {
	gorm.Model
	Name         string `gorm:"unique; not null"`
	Abbreviation string `gorm:"unique; not null"`
}

type Movement struct {
	gorm.Model
	Name           string `gorm:"unique; not null"`
	Abbreviation   string `gorm:"unique; not null"`
	ImageUrl       string
	MovementTypeID uint
}

type createMovementTypeRequest struct {
	Name         string
	Abbreviation string
}

func CreateMovementType(c *gin.Context) {
	db := common.GetDB()

	var r createMovementTypeRequest

	if err := c.BindJSON(&r); err != nil {
		fmt.Errorf("Unable to handle request")
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
