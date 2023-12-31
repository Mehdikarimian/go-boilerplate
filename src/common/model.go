package common

import (
	"github.com/go-gorp/gorp"
	_redis "github.com/go-redis/redis/v7"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type ModelConstructor struct {
	Collection *mongo.Collection
	Repository *gorp.DbMap
	Redis      *_redis.Client
	Gorm       *gorm.DB
}
