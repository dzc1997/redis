package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
)

var redisClient *redis.Client

func initRedis() error {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "0.0.0.0:6379",
		Password: "19971002",
		DB:       0,
	})
	_, err := redisClient.Ping(context.Background()).Result()
	return err
}

func main() {
	err := initRedis()
	if err != nil {
		fmt.Printf("connet redis failed, err: %v\n", err)
		return
	}
	fmt.Println("redis connect success!")
}
