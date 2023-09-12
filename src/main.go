package main

import (
	"fmt"
	"go-boilerplate/src/config"
	"go-boilerplate/src/controllers"
	"go-boilerplate/src/core/db"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db.InitRedis(1)
	db.InitMongoDB()
	db.InitPostgresDB()

	// register controllers
	controllers.UsersController(r)
	controllers.ArticlesController(r)

	r.Use(gzip.Gzip(gzip.DefaultCompression))

	// running
	r.Run(fmt.Sprintf(":%s", config.LoadConfig("PORT")))
}
