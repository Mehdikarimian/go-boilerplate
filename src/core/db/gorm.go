package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var gormdb *gorm.DB

func connectGorm() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=localhost port=5432 user=postgres dbname=go password=postgres sslmode=disable")

	fmt.Println("connected to " + dsn)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func InitGorm() *gorm.DB {
	if gormdb != nil {
		return gormdb
	}

	var err error
	gormdb, err = connectGorm()

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return gormdb
}

func GetGorm() *gorm.DB {

	return gormdb
}
