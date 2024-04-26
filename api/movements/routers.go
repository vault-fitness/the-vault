package movements

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"vaultfitness.io/api/common"
)

func MovementsRegister(router *gin.RouterGroup) {

	router.GET("/type", MovementTypesList)
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

func MovementTypesList(c *gin.Context) {
	db := common.GetDB()

	var mt []MovementType
	db.Raw("SELECT id, name, abbreviation FROM movement_types").Scan(&mt)

	c.JSON(http.StatusOK, mt)
}
