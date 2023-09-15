package utils

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetBody[ValidatorStruct any](ctx *gin.Context) ValidatorStruct {
	var body ValidatorStruct

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	return body
}

func GetParam[ValidatorStruct any](ctx *gin.Context) ValidatorStruct {
	var param ValidatorStruct

	if err := ctx.ShouldBindUri(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err})
	}

	return param
}

func GetQueryInt(ctx *gin.Context, Key string, defaultValue int) int {
	query := ctx.Request.URL.Query()

	value, err := strconv.Atoi(query.Get(Key))

	if err != nil {
		value = defaultValue
	}

	return value
}

func GetQueryString(ctx *gin.Context, Key string, defaultValue string) string {
	value := ctx.Request.URL.Query().Get(Key)

	return value
}
