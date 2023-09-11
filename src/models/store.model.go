package models

import (
	"go-boilerplate/src/common"
	"go-boilerplate/src/core"
)

func StoresModel() *BaseModel {
	mod := &BaseModel{
		ModelConstructor: &common.ModelConstructor{
			Collection: core.InitMongoDB().Collection("testStoresCollection"),
		},
	}

	return mod
}

// models definitions
type Store struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Type     string `form:"type" json:"type" binding:"required"`
	Category string `form:"category" json:"category" binding:"required"`
}

// models methods
func (mod *BaseModel) GetOneStore() {

}

func (mod *BaseModel) GetAllStores() {

}

func (mod *BaseModel) UpdateStore() {

}

func (mod *BaseModel) CreateStore() {

}
