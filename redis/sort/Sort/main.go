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
	SortList()
	SortHash()
}

func SortList() {
	res := conn.Sort(ctx, "sort-input", &redis.Sort{Order: "ASC"}).Val()
	fmt.Println("res = ", res)
}

func SortHash() {
	res := conn.Sort(ctx, "sort-input", &redis.Sort{Alpha: true}).Val()
	res = conn.Sort(ctx, "sort-input", &redis.Sort{By: "d-*->field"}).Val()
	res = conn.Sort(ctx, "sort-input", &redis.Sort{By: "d-*->field", Get: []string{"d-*->field"}}).Val()
	fmt.Println("res = ", res)
}
