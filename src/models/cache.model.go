package models

import (
	"go-boilerplate/src/common"
	"go-boilerplate/src/core/db"
	"time"
)

func CacheModel() *BaseModel {
	mod := &BaseModel{
		ModelConstructor: &common.ModelConstructor{
			Redis: db.GetRedis(),
		},
	}

	return mod
}

func (mod *BaseModel) GetCache(key string) string {
	val, err := mod.Redis.Get(key).Result()

	if err != nil {
		return ""
	}

	return val
}

func (mod *BaseModel) SetCache(key string, value interface{}, ttl time.Duration) {
	err := mod.Redis.Set(key, value, ttl).Err()
	if err != nil {
		panic(err)
	}
}
