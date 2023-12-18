package sms

import "context"

type Service interface {
	SendTencent(ctx context.Context, tpl string, args []string, numbers ...string) error
	SendAliyun(ctx context.Context, tpl string, args []string, numbers ...string) error
	SendAliyunV1(ctx context.Context, tpl string, args []NameArg, numbers ...string) error
}

type NameArg struct {
	Name string
	Val  string
}
