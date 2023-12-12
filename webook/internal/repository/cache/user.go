package cache

import (
	"basic-go/webook/internal/domain"
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

//type Cache interface {
//	GETUER(ctx context.Context, id int64) (domain.User, error)
//	//读取文章
//	GETArticle(ctx context.Context, aid int64)
//}
//
//type CacheV1 interface {
//	//中间件团队去做
//	GET(ctx context.Context, key string) (any, error)
//}

var ErrNotExist = redis.Nil

type UserCache struct {
	//cache CacheV1
	//传单机redis可以
	//传cluster 的 redis也可以
	client redis.Cmdable
	//超时
	expiration time.Duration
}

// NewUserCache
// A用到了 B ，B一定是接口
// A用到了 B ，B一定是A的字段
// A用到了 B ，A绝对不初始化B，而是从外部注入
func NewUserCache(client redis.Cmdable) *UserCache {
	return &UserCache{
		client:     client,
		expiration: time.Minute * 15,
	}
}

// 只要error为niu，就认为缓存有数据
// 如果没有数据，返回一个特定的error
func (cache *UserCache) Get(ctx context.Context, id int64) (domain.User, error) {
	key := cache.key(id)
	//数据不存在，err=redis.Nil
	val, err := cache.client.Get(ctx, key).Bytes()
	if err != nil {
		return domain.User{}, err
	}
	var u domain.User
	err = json.Unmarshal(val, &u)
	return u, err
}

func (cache *UserCache) Set(ctx context.Context, u domain.User) error {
	val, err := json.Marshal(u)
	if err != nil {
		return err
	}
	key := cache.key(u.Id)
	return cache.client.Set(ctx, key, val, cache.expiration).Err()
}

func (cache *UserCache) key(id int64) string {
	return fmt.Sprintf("user:info:%d", id)

}
