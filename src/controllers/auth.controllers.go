package controllers

import (
	"go-boilerplate/src/common"
	"go-boilerplate/src/middleware"
	"go-boilerplate/src/models"
	"go-boilerplate/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthController(r *gin.Engine) *BaseController {
	ctr := &BaseController{
		ControllerConstructor: &common.ControllerConstructor{
			// Collection: core.InitMongoDB().Collection("testUsersCollection"),
		},
	}

	ctr.AuthRoutes(r)

	return ctr
}

// @BasePath /users
func (ctr *BaseController) AuthRoutes(r *gin.Engine) {
	api := r.Group("/auth")
	{
		api.POST("/login", func(ctx *gin.Context) {
			login(ctx, ctr)
		})
	}
	api.Use(middleware.JwtAuthMiddleware())
	{
		api.GET("/profile", func(ctx *gin.Context) {
			profile(ctx, ctr)
		})
	}
}

// @Summary      Login
// @Tags         auth
// @Produce      json
// @Success      200  {array}  models.User
// @Router       /auth/login [post]
// @Param request body models.LoginInput true "body"
func login(ctx *gin.Context, ctr *BaseController) {
	input := utils.GetBody[models.LoginInput](ctx)

	token, err := models.AuthModel().LoginCheck(input.Username, input.Password)

	if err != nil || token == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "The username or password is not correct"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

// @Summary      Profile
// @Tags         auth
// @Produce      json
// @Success      200  {array}  models.User
// @Router       /auth/profile [get]
// @Security jwtTokenAuth
func profile(ctx *gin.Context, ctr *BaseController) {
	user, err := models.AuthModel().CurrentUser(ctx)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err)
	}

	ctx.JSON(http.StatusOK, user)
}
