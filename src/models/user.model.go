package models

import (
	"context"
	"fmt"
	"go-boilerplate/src/common"
	"go-boilerplate/src/core/db"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func MigrateUsers() {
	db.GetGorm().AutoMigrate(&User{})
}

func UsersModel() *BaseModel {
	mod := &BaseModel{
		ModelConstructor: &common.ModelConstructor{
			Gorm: db.GetGorm(),
		},
	}

	return mod
}

// models definitions

type User struct {
	ID        uint `gorm:"autoIncrement,primaryKey"`
	Name      string
	Email     *string
	LastName  uint8
	FirstName uint8
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UsersFindParam struct {
	ID uint `uri:"id" binding:"required"`
}

// models methods
func (mod *BaseModel) GetOneUser(param User) mongo.SingleResult {
	result := mod.Collection.FindOne(context.TODO(), bson.M{"_id": param.ID})

	return *result
}

func (mod *BaseModel) GetAllUsers() *User {
	user := &User{}

	result := mod.Gorm.Table("users").Select("name", "age").Scan(&user)

	if result.Error != nil {
		fmt.Println(result.Error)
		return nil
	}

	return user
}

func (mod *BaseModel) UpdateUser() {

}

func (mod *BaseModel) CreateUser(body User) {
	mod.Collection.InsertOne(context.TODO(), bson.M{"name": body.Name, "email": body.Email, "lastName": body.LastName})

	return
}
