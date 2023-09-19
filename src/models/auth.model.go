package models

import (
	"fmt"
	"go-boilerplate/src/common"
	"go-boilerplate/src/config"
	"go-boilerplate/src/core/db"
	"go-boilerplate/src/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func AuthModel() *BaseModel {
	mod := &BaseModel{
		ModelConstructor: &common.ModelConstructor{
			Gorm: db.GetGorm(),
		},
	}

	return mod
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (mod *BaseModel) CurrentUser(ctx *gin.Context) (User, error) {
	var user User

	user_id, err := utils.ExtractTokenID(ctx)
	fmt.Println(user)

	if err != nil {
		return user, err
	}

	user = UsersModel().GetOneUser(user_id)

	if err != nil {
		return user, err
	}

	return user, nil
}

func GenerateToken(user User) (string, error) {

	tokenLifespan, err := strconv.Atoi(config.LoadConfig("TOKEN_HOUR_LIFESPAN"))

	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(tokenLifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.LoadConfig("API_SECRET")))

}

func (mod *BaseModel) LoginCheck(username, password string) (string, error) {
	var err error

	user := User{}

	err = mod.Gorm.Limit(1).Where("username=?", username).Find(&user).Error

	if err != nil || user.ID == 0 {
		return "", err
	}

	err = VerifyPassword(password, user.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := GenerateToken(user)

	if err != nil {
		return "", err
	}

	return token, nil
}
