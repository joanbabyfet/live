package config

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client
//Go 标准库 context.Context, 用来传递超时、取消信号、请求链路信息, 通常给 Redis、MySQL、HTTP Client 等使用
var Ctx = context.Background()

func InitRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   6,
	})
}