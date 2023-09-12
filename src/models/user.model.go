package models

import (
	"context"
	"go-boilerplate/src/common"
	"go-boilerplate/src/core/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UsersModel() *BaseModel {
	mod := &BaseModel{
		ModelConstructor: &common.ModelConstructor{
			Collection: db.GetMongoDb().Collection("testUsersCollection"),
		},
	}

	return mod
}

// models definitions

type User struct {
	ID       string `form:"_id" json:"_id"`
	Name     string `form:"name" json:"name" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required,min=8"`
}

type UsersFindParam struct {
	ID string `uri:"id" binding:"required"`
}

// models methods
func (mod *BaseModel) GetOneUser(param User) mongo.SingleResult {
	result := mod.Collection.FindOne(context.TODO(), bson.M{"_id": param.ID})

	return *result
}

func (mod *BaseModel) GetAllUsers() {

}

func (mod *BaseModel) UpdateUser() {

}

func (mod *BaseModel) CreateUser(body User) {
	mod.Collection.InsertOne(context.TODO(), bson.M{"name": body.Name, "email": body.Email, "password": body.Password})

	return
}
