package main

import (
	"fmt"
	"go-boilerplate/src/config"
	"go-boilerplate/src/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// register controllers
	controllers.UsersController(r)
	controllers.StoresController(r)

	// running
	r.Run(fmt.Sprintf(":%s", config.LoadConfig("PORT")))
}
