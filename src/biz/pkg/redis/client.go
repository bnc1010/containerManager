package redis

import (
	"context"
	"fmt"
	// "time"
	viper "github.com/bnc1010/containerManager/biz/pkg/viper"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	redisv8 "github.com/go-redis/redis/v8"
)

var (
	Client *redisv8.Client
	config *viper.Redis
)

// InitRedis 初始化Redis
func InitRedis() bool {
	config = viper.Conf.Redis
	return initRedis(context.Background(), &Client)
}

// initRedis 初始化Redis impl
func initRedis(ctx context.Context, client **redisv8.Client) bool {
	rdb := redisv8.NewClient(&redisv8.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.Db,
	})
	if rdb == nil {
		hlog.CtxFatalf(ctx, "[Redis] Init Failed")
		return false
	}
	*client = rdb
	pingMsg := fmt.Sprintf("%s", Client.Ping(ctx))
	hlog.CtxInfof(ctx, "[Redis] PING: %s\n", pingMsg)
	if pingMsg != "ping: PONG" {
		hlog.CtxInfof(ctx, "[Redis] connect timeout")
		return false
	}
	return true
}


// GetIncrId 获取Redis计数器
func GetIncrId(ctx context.Context, key string) int64 {
	id, err := Client.Incr(ctx, key).Result()
	if err != nil {
		hlog.CtxErrorf(ctx, "[Redis] GetIncrId failed, key: %v, err: %v", key, err)
		return 0
	}
	return id
}

// GetValue 获取redis value
func GetValue(ctx context.Context, key string) string {
	res, err := Client.Get(ctx, key).Result()
	if err != nil {
		hlog.CtxErrorf(ctx, "[Redis] GetIncrId failed, key: %v, err: %v", key, err)
		return ""
	}
	return res
}
