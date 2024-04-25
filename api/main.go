package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"vaultfitness.io/api/common"
	"vaultfitness.io/api/movements"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}
	common.Init()

	r := gin.Default()

	v1 := r.Group("/v1")
	movements.MovementsRegister(v1.Group("/movements"))

	r.Run()
}
