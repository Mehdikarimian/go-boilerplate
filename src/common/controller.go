package common

import (
	"github.com/gin-gonic/gin"
	"github.com/go-gorp/gorp"
	"go.mongodb.org/mongo-driver/mongo"
)

type ControllerConstructor struct {
	Collection *mongo.Collection
	Repository *gorp.DbMap
}

type ControllerInterface interface {
	RegisterRoutes(*gin.Engine)
}
