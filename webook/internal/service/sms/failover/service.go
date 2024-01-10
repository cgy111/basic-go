package failover

import (
	"basic-go/webook/internal/service/sms"
	"context"
	"errors"
	"log"
)

type FailoverSMSService struct {
	svcs []sms.Service
}

func (f FailoverSMSService) Send(ctx context.Context, tpl string, args []string, numbers ...string) error {
	for _, svc := range f.svcs {
		err := svc.Send(ctx, tpl, args, numbers...)
		//发生成功
		if err == nil {
			return nil
		}
		//正常这里，输出日志
		//要做好监控
		log.Println(err)
	}
	return errors.New("全部服务商都失败了")
}

func NewFailoverSMSService(svcs []sms.Service) sms.Service {
	return &FailoverSMSService{
		svcs: svcs,
	}
}
