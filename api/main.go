package main

import (
	"github.com/gin-gonic/gin"
	"vaultfitness.io/api/common"
	"vaultfitness.io/api/movements"
)

func main() {
	// err := godotenv.Load()

	// if err != nil {
	// 	panic("Error loading .env file")
	// }

	r := gin.Default()

	v1 := r.Group("/api/v1")
	movements.MovementsRegister(v1.Group("/movements"))

	r.Run()

	common.Init()
}
