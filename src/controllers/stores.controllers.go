package controllers

import (
	"go-boilerplate/src/common"

	"github.com/gin-gonic/gin"
)

func StoresController(r *gin.Engine) *BaseController {
	ctr := &BaseController{
		ControllerConstructor: &common.ControllerConstructor{
			// Collection: core.InitMongoDB().Collection("testStoresCollection"),
		},
	}

	ctr.StoresRoutes(r)

	return ctr
}

func (ctr *BaseController) StoresRoutes(r *gin.Engine) {
	api := r.Group("/stores")
	{
		api.GET("/", func(ctx *gin.Context) {

		})

		api.POST("/", func(ctx *gin.Context) {

		})

		api.PUT("/:id", func(ctx *gin.Context) {

		})
	}
}
