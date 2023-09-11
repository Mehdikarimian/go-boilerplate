package common

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type ModelConstructor struct {
	Collection *mongo.Collection
}
