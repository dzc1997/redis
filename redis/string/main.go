package main

import (
	"context"
	"fmt"
	"gitee.com/wedone/redis_course/redis/redisConn"
)

var conn = redisConn.ConnectRedis()

var ctx = context.Background()

func main() {
	connRedis()
	Get()
	Set()
	Incr()
	IncrBy()
	Decr()
	DecrBy()
}

func connRedis() {
	fmt.Println(ctx, "conn = ", conn) // Redis<62.234.11.179:6379 db:15>
}

func Set() {
	res := conn.Set(ctx,"num", 9, 0).Val()
	fmt.Println("res = ", res)
}

func Get() {
	res := conn.Get(ctx,"num")
	fmt.Println("num=", res)
}

func Incr() { // incr自增
	res := conn.Incr(ctx,"num").Val()
	fmt.Println("res = ", res)
}

func IncrBy() {
	res := conn.IncrBy(ctx,"num", 3)
	fmt.Println("res = ", res)
}

func Decr() {
	res := conn.Decr(ctx,"num")
	fmt.Println("res = ", res)
}

func DecrBy() {
	res := conn.DecrBy(ctx,"num", 5)
	fmt.Println("res = ", res)
}
