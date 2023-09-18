package db

import (
	"fmt"
	"go-boilerplate/src/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var gormdb *gorm.DB

func connectGorm() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", config.LoadConfig("POSTGRES_HOST"), config.LoadConfig("POSTGRES_PORT"), config.LoadConfig("POSTGRES_USER"), config.LoadConfig("POSTGRES_DB_NAME"), config.LoadConfig("POSTGRES_PASSWORD"))

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
