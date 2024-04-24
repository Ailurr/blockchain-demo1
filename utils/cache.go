package utils

import (
	"context"
	"demo1/global"
	"github.com/redis/go-redis/v9"
	"math/rand"
	"time"
)

func GetCache(key string) ([]byte, error) {
	ctx := context.Background()
	stringCmd := global.Rdc.Get(ctx, key)
	if stringCmd.Err() != nil {
		if stringCmd.Err() == redis.Nil {
			return nil, stringCmd.Err()
		}
	}
	bytes, err := stringCmd.Bytes()
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func SetCache(key string, value string) error {
	ctx := context.Background()
	expiredMinute := rand.Intn(15) + 10
	//expiredMinute := rand.Intn(1) + 1
	statusCmd := global.Rdc.Set(ctx, key, value, time.Minute*time.Duration(expiredMinute))
	return statusCmd.Err()
}
