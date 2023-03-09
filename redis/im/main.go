package main

import (
	"context"
	"fmt"
	"gitee.com/wedone/redis_course/redis/redisConn"
	"github.com/go-redis/redis/v9"
	uuid "github.com/satori/go.uuid"
	"log"
	"math"
	"strconv"
	"strings"
	"time"
)

var ctx = context.Background()
var conn = redisConn.ConnectRedis()

func main() {
	user := CreateUser(ctx, "123", "shuwen")
	fmt.Println("user = ", user)
}

func CreateUser(ctx context.Context, login, name string) string {
	llogin := strings.ToLower(login)
	lock := AcquireLockWithTimeout(ctx, "user:"+llogin, 10, 10) // 加锁函数对小写的用户名进行加锁
	defer ReleaseLock(ctx, "user:"+llogin, lock)
	if lock == "" {
		return ""
	}
	if conn.HGet(ctx, "users:", llogin).Val() != "" {
		return ""
	}
	id := conn.Incr(ctx, "user:id:").Val()
	pipeline := conn.TxPipeline()
	pipeline.HSet(ctx, "users:", llogin, id)
	pipeline.HMSet(ctx, fmt.Sprintf("user:%s", strconv.Itoa(int(id))), "login", login, "id", id, "name", name, "followers", 0, "following", 0, "posts", 0, "sigup", time.Now().UnixNano())
	if _, err := pipeline.Exec(ctx); err != nil {
		log.Println("pipeline err in CreateUser:", err)
		return ""
	}
	return strconv.Itoa(int(id))
}

func AcquireLockWithTimeout(ctx context.Context, lockname string, acquireTimeout, lockTimeout float64) string {
	identifier := uuid.NewV4().String()
	lockname = "lock" + lockname
	finallLockTimeout := math.Ceil(lockTimeout)

	end := time.Now().UnixNano() + int64(acquireTimeout*1e9)
	for time.Now().UnixNano() < end {
		if conn.SetNX(ctx, lockname, identifier, 0).Val() {
			conn.Expire(ctx, lockname, time.Duration(finallLockTimeout)*time.Second)
			return identifier
		} else if conn.TTL(ctx, lockname).Val() < 0 {
			conn.Expire(ctx, lockname, time.Duration(finallLockTimeout)*time.Second)
		}
		time.Sleep(10 * time.Millisecond)
	}
	return ""
}

func ReleaseLock(ctx context.Context, lockname, identifier string) bool {
	lockname = "lock:" + lockname
	var flag = true
	for flag {
		err := conn.Watch(ctx, func(tx *redis.Tx) error {
			pipe := tx.TxPipeline()
			fmt.Println(pipe)
			if tx.Get(ctx, lockname).Val() == identifier {
				pipe.Del(ctx, lockname)
				if _, err := pipe.Exec(ctx); err != nil {
					return err
				}
				flag = true
				return nil
			}
			tx.Unwatch(ctx)
			flag = false
			return nil
		})
		if err != nil {
			log.Println("watch failed in ReleaseLock,err is:", err)
			return false
		}
	}
	return true
}
