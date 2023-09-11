package utils

import (
	"net/http"

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
		ctx.JSON(400, gin.H{"msg": err})
	}

	return param
}
