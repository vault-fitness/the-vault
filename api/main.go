package main

import (
	"github.com/joho/godotenv"
	"vaultfitness.io/api/common"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}

	common.Init()
}
