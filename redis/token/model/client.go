package model

import (
	"context"
	"github.com/go-redis/redis/v9"
	"time"
)

var ctx context.Context

type Client struct {
	Conn *redis.Client
}

// CheckToken 尝试获取并返回令牌对应的用户
func (r *Client) CheckToken(token string) string {
	return r.Conn.HGet(ctx, "login:", token).Val()
}

// UpdateToken 更新令牌
func (r *Client) UpdateToken(token, user, item string) {
	timestamp := time.Now().Unix()
	r.Conn.HSet(ctx, "login:", token, user)
	r.Conn.ZAdd(ctx, "recent", redis.Z{Score: float64(timestamp), Member: token})
	if item != "" {
		//r.Conn.ZAdd(ctx, "viewed:"+token, item, float64(timestamp))
		r.Conn.ZAdd(ctx, "viewed"+token, redis.Z{
			Score:  float64(timestamp),
			Member: item,
		})
		r.Conn.ZRemRangeByRank(ctx, "viewed:"+token, 0, -26)
	}
}
