package ratelimit

import "context"

type Limiter interface {
	//Limit 有没有触发限流,key就是限流对象
	//bool 是否触发限流,true 触发限流,false 没有触发限流
	//err 限流本身有没有错误
	Limit(ctx context.Context, key string) (bool, error)
}
