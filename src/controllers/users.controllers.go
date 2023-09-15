package controllers

import (
	"fmt"
	"go-boilerplate/src/common"
	"go-boilerplate/src/models"
	"go-boilerplate/src/utils"
	"net/http"

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
		api.DELETE("/:id", func(ctx *gin.Context) {
			deleteUser(ctx, ctr)
		})
	}
}

// @Summary      List Of users
// @Tags         users
// @Produce      json
// @Success      200  {array}  models.User
// @Router       /users [get]
// @Param        limit    query    int  false  "limit"  Format(limit)
// @Param        skip    query    int  false  "skip"  Format(skip)
// @Param        search    query    string  false  "search"  Format(search)
func getUsers(ctx *gin.Context, ctr *BaseController) {
	limit := utils.GetQueryInt(ctx, "limit", 20)
	skip := utils.GetQueryInt(ctx, "skip", 0)
	search := utils.GetQueryString(ctx, "search", "")

	users, count := models.UsersModel().GetAllUsers(limit, skip, search)

	response := models.UsersResponse{
		Users: users,
		Count: count,
	}

	ctx.JSON(http.StatusOK, response)
}

// @Summary      Get An User
// @Tags         users
// @Produce      json
// @Success      200  {object}  models.User
// @Router       /users/{id} [get]
// @Param        id    path    int  false  "id"  Format(id)
func getUser(ctx *gin.Context, ctr *BaseController) {
	param := utils.GetParam[models.UsersFindParam](ctx)

	results := models.UsersModel().GetOneUser(models.UsersFindParam{ID: param.ID})

	if results == (models.User{}) {
		ctx.JSON(http.StatusNotFound, "User Not Found!")
		return
	}
	ctx.JSON(http.StatusOK, results)
}

// @Summary      Create An User
// @Description  Create An User.
// @Tags         users
// @Produce      json
// @Success      201  {object}  models.User
// @Router       /users [post]
// @Param request body models.CreateUserForm true "body"
func createUser(ctx *gin.Context, ctr *BaseController) {
	body := utils.GetBody[models.User](ctx)

	result := models.UsersModel().CreateUser(body)

	ctx.JSON(http.StatusCreated, result)
}

// @Summary      Update An User
// @Description  Update An User.
// @Tags         users
// @Produce      json
// @Success      201  {object}  models.User
// @Router       /users/{id} [put]
// @Param        id    path    int  false  "id"  Format(id)
// @Param request body models.CreateUserForm true "body"
func updateUser(ctx *gin.Context, ctr *BaseController) {
	param := utils.GetParam[models.UsersFindParam](ctx)
	body := utils.GetBody[models.User](ctx)

	fmt.Println(param)
	results := models.UsersModel().UpdateUser(models.UsersFindParam{ID: param.ID}, body)

	ctx.JSON(http.StatusCreated, results)
}

// @Summary      Delete An User
// @Description  Delete An User.
// @Tags         users
// @Produce      json
// @Success      200  {object}  models.User
// @Router       /users/{id} [delete]
// @Param        id    path    int  false  "id"  Format(id)
func deleteUser(ctx *gin.Context, ctr *BaseController) {
	param := utils.GetParam[models.UsersFindParam](ctx)

	results := models.UsersModel().DeleteUser(models.UsersFindParam{ID: param.ID})

	ctx.JSON(http.StatusOK, results)
}
