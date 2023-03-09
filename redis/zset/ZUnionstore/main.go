package main

import (
	"context"
	"fmt"
	"gitee.com/wedone/redis_course/redis/redisConn"
	"github.com/go-redis/redis/v9"
)

var conn = redisConn.ConnectRedis()
var ctx = context.Background()

func main() {
	ZUnionStore()
}

func ZUnionStore() {
	res := conn.ZUnionStore(ctx,"zset-u", &redis.ZStore{Aggregate: "min", Keys: []string{"zset-1", "zset-2"}}).Val()
	fmt.Println("res = ", res)
}
