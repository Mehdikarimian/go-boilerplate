package middleware

import (
	"fmt"
	"net/http"

	"go-boilerplate/src/utils"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		err := utils.TokenValid(ctx)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"Unauthorized": "Authentication required"})
			fmt.Println(err)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
