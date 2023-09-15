package models

import (
	"fmt"
	"go-boilerplate/src/common"
	"go-boilerplate/src/core/db"
	"time"
)

func UsersModel() *BaseModel {
	mod := &BaseModel{
		ModelConstructor: &common.ModelConstructor{
			Gorm: db.GetGorm(),
		},
	}

	return mod
}

func MigrateUsers() {
	db.GetGorm().AutoMigrate(&User{})
}

// models definitions

type User struct {
	ID        uint   `gorm:"autoIncrement,primaryKey"`
	Username  string `gorm:"not null,index,unique"`
	Email     *string
	FirstName string `gorm:"default:''"`
	LastName  string `gorm:"default:''"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateUserForm struct {
	Username  string `form:"Username" json:"Username" binding:"required"`
	Email     string `form:"Email" json:"Email" binding:"required"`
	FirstName string `form:"FirstName" json:"FirstName" binding:"required"`
	LastName  string `form:"LastName" json:"LastName" binding:"required"`
}

type UsersResponse struct {
	Users []User `json:"users"`
	Count int    `json:"count"`
}
type UsersFindParam struct {
	ID uint `uri:"id" binding:"required"`
}

// models methods
func (mod *BaseModel) GetOneUser(param UsersFindParam) User {
	var user User

	result := mod.Gorm.Limit(1).Where("id = ?", param.ID).Find(&user)

	if result.Error != nil {
		fmt.Println(result.Error)
		return user
	}

	return user
}

func (mod *BaseModel) GetAllUsers(limit int, skip int, search string) ([]User, int) {
	var users []User
	var count int

	search = "%" + search + "%"

	result := mod.Gorm.Limit(limit).Offset(skip).Where("email ilike ? OR username ilike ? OR first_name ilike ? OR last_name ilike ?", search, search, search, search).Find(&users)

	mod.Gorm.Raw("SELECT COUNT(id) FROM users WHERE email ilike ? OR username ilike ? OR first_name ilike ? OR last_name ilike ?", search, search, search, search).Scan(&count)

	if result.Error != nil {
		fmt.Println(result.Error)
		return nil, 0
	}

	return users, count
}

func (mod *BaseModel) UpdateUser(param UsersFindParam, body User) User {
	body.ID = param.ID

	result := mod.Gorm.Save(&body)

	if result.Error != nil {
		fmt.Println(result.Error)
	}

	return body
}

func (mod *BaseModel) CreateUser(body User) User {
	result := mod.Gorm.Create(&body)

	if result.Error != nil {
		fmt.Println(result.Error)
	}

	return body
}

func (mod *BaseModel) DeleteUser(param UsersFindParam) bool {
	var user User
	user.ID = param.ID

	result := mod.Gorm.Delete(&user)

	fmt.Println(result)
	if result.Error != nil || result.RowsAffected == 0 {
		return false
	}

	return true
}
