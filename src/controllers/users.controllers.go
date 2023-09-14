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

// @BasePath /users
func (ctr *BaseController) UsersRoutes(r *gin.Engine) {
	api := r.Group("/users")
	{
		api.GET("/", func(ctx *gin.Context) {
			getUsers(ctx, ctr)
		})
		api.GET("/:id", func(ctx *gin.Context) {
			getUser(ctx, ctr)
		})
		api.POST("/", func(ctx *gin.Context) {
			createUser(ctx, ctr)
		})
		api.PUT("/:id", func(ctx *gin.Context) {
			updateUser(ctx, ctr)
		})
	}
}

// @Summary      List Of users
// @Tags         users
// @Produce      json
// @Success      200  {array}  models.User
// @Router       /users [get]
func getUsers(ctx *gin.Context, ctr *BaseController) {
	results := models.UsersModel().GetAllUsers()

	ctx.JSON(200, results)
}

// @Summary      Get An User
// @Tags         users
// @Produce      json
// @Success      200  {object}  models.User
// @Router       /users/:id [get]
func getUser(ctx *gin.Context, ctr *BaseController) {
	param := utils.GetParam[models.UsersFindParam](ctx)

	results := models.UsersModel().GetOneUser(models.User{ID: param.ID})

	ctx.JSON(200, results)
}

// @Summary      Create An User
// @Description  Create An User.
// @Tags         users
// @Produce      json
// @Success      200  {object}  models.User
// @Router       /users [post]
func createUser(ctx *gin.Context, ctr *BaseController) {
	body := utils.GetBody[models.User](ctx)

	models.UsersModel().CreateUser(body)

	ctx.JSON(200, "User Created")
}

// @Summary      Update An User
// @Description  Update An User.
// @Tags         users
// @Produce      json
// @Success      200  {object}  models.User
// @Router       /users [put]
func updateUser(ctx *gin.Context, ctr *BaseController) {
	param := utils.GetParam[models.UsersFindParam](ctx)

	results := models.UsersModel().GetOneUser(models.User{ID: param.ID})

	ctx.JSON(200, results)
}

// @Summary      Delete An User
// @Description  Delete An User.
// @Tags         users
// @Produce      json
// @Success      200  {object}  models.User
// @Router       /users [delete]
func deleteUser(ctx *gin.Context, ctr *BaseController) {
	param := utils.GetParam[models.UsersFindParam](ctx)

	results := models.UsersModel().GetOneUser(models.User{ID: param.ID})

	ctx.JSON(200, results)
}
