package retryable

//
//import (
//	"basic-go/webook/internal/service/sms"
//	"context"
//)
//
//// 小心并发问题
//// Service 重试机制
//type Service struct {
//	svc      sms.Service
//	retryCnt int
//}
//
//func (s Service) SendTencent(ctx context.Context, tpl string, args []string, numbers ...string) error {
//	err := s.svc.SendTencent(ctx, tpl, args, numbers...)
//	for err != nil && s.retryCnt < 10 {
//		err = s.svc.SendTencent(ctx, tpl, args, numbers...)
//		s.retryCnt++
//	}
//	return err
//}
