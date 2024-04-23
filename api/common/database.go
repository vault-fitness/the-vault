package common

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

// Opening a database and save the reference to `Database` struct.
func Init() *gorm.DB {
	connection_string := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(connection_string), &gorm.Config{})

	if err != nil {
		fmt.Println("db err: (Init) ", err)
	}

	DB = db

	return DB
}

// Using this function to get a connection, you can create your connection pool here.
func GetDB() *gorm.DB {
	return DB
}
