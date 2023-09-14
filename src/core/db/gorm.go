package db

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var gormdb *gorm.DB

type User struct {
	ID        uint `gorm:"autoIncrement,primaryKey"`
	Name      string
	Email     *string
	Age       uint8
	CreatedAt time.Time
	UpdatedAt time.Time
}

func connectGorm() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=localhost port=5432 user=postgres dbname=go password=postgres sslmode=disable")

	fmt.Println("connected to " + dsn)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func InitGorm() *gorm.DB {
	gormdb, err := connectGorm()

	if err != nil {
		log.Fatal(err)
	}

	return gormdb
}

func GetGorm() *gorm.DB {

	return gormdb
}
