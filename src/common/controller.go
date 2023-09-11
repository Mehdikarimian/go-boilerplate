package common

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type ControllerConstructor struct {
	Collection *mongo.Collection
}

type ControllerInterface interface {
	RegisterRoutes(*gin.Engine)
}
