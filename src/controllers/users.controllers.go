package controllers

import (
	"go-boilerplate/src/common"
	"go-boilerplate/src/models"
	"go-boilerplate/src/utils"

	"github.com/gin-gonic/gin"
)

func UsersController(r *gin.Engine) *BaseController {
	ctr := &BaseController{
		ControllerConstructor: &common.ControllerConstructor{
			// Collection: core.InitMongoDB().Collection("testUsersCollection"),
		},
	}

	ctr.UsersRoutes(r)

	return ctr
}

func (ctr *BaseController) UsersRoutes(r *gin.Engine) {
	api := r.Group("/users")
	{
		api.GET("/", func(ctx *gin.Context) {
			results := models.UsersModel().GetAllUsers()

			ctx.JSON(200, results)
		})

		api.GET("/:id", func(ctx *gin.Context) {
			param := utils.GetParam[models.UsersFindParam](ctx)

			results := models.UsersModel().GetOneUser(models.User{ID: param.ID})

			ctx.JSON(200, results)
		})

		api.POST("/", func(ctx *gin.Context) {
			body := utils.GetBody[models.User](ctx)

			models.UsersModel().CreateUser(body)

			ctx.JSON(200, "User Created")
		})

		api.PUT("/:id", func(ctx *gin.Context) {
			param := utils.GetParam[models.UsersFindParam](ctx)

			results := models.UsersModel().GetOneUser(models.User{ID: param.ID})

			ctx.JSON(200, results)
		})
	}
}
