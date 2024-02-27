package ioc

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

func InitRedis() redis.Cmdable {
	//addr := viper.GetString("redis.addr")

	type Config struct {
		Addr string `yaml:"addr"`
	}
	var cfg Config

	viper.UnmarshalKey("redis", &cfg)

	redisClient := redis.NewClient(&redis.Options{
		//Addr: addr,
		Addr: cfg.Addr,
	})
	return redisClient
}

//func NewRateLimiter() r {
//
//}
