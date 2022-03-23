package redisx

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
	"younghe/config"
)

var Redis *redis.Client

func Setup() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     config.Config.Redis.Address,
		Password: config.Config.Redis.Password,
		DB:       config.Config.Redis.DB,
	})
	if err := Redis.Ping(context.TODO()).Err(); err != nil {
		panic(err)
	}
}

func TryLock(key, token string, expiration uint) (bool, error) {
	return Redis.SetNX(context.TODO(), key, token, time.Second*time.Duration(expiration)).Result()
}
