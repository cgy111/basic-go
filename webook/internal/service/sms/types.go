package sms

import "context"

type Service interface {
	Send(ctx context.Context, tpl string, args []string, numbers ...string) error
	//SendAliyun(ctx context.Context, tpl string, args []string, numbers ...string) error
	//SendAliyunV1(ctx context.Context, tpl string, args []NameArg, numbers ...string) error
	////缺点：调用者需要知道实现者需要什么参数，是[]string,还是map[string]string，不推荐
	//SendAliyunV2(ctx context.Context, tpl string, args any, numbers ...string) error
}

//type NameArg struct {
//	Name string
//	Val  string
//}
