package main

import (
	"context"
	"fmt"
	"gitee.com/wedone/redis_course/redis/redisConn"
)

var conn = redisConn.ConnectRedis()

var ctx = context.Background()

func main() {
	SetBit()
	SetBit2()
	Get()
}

func SetBit() {
	res := conn.SetBit(ctx, "another-key", 2, 1).Val() // 将第2个二进制位(offset)设置为1
	fmt.Println("res = ", res)
}

func SetBit2() {
	res := conn.SetBit(ctx, "another-key", 7, 1).Val() // 将第7个二进制位(offset)设置为1
	fmt.Println("res = ", res)
}

func Get() {
	str := conn.Get(ctx, "another-key").Val() // 键的值将会变为！=> ASCII码为33的字符
	fmt.Println("str = ", str)
}
