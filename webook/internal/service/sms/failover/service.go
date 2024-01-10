package failover

import (
	"basic-go/webook/internal/service/sms"
	"context"
	"errors"
	"log"
	"sync/atomic"
)

type FailoverSMSService struct {
	svcs []sms.Service

	idx uint64
}

func (f *FailoverSMSService) Send(ctx context.Context, tpl string, args []string, numbers ...string) error {
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

func (f *FailoverSMSService) SendV1(ctx context.Context, tpl string, args []string, numbers ...string) error {
	//取下一个节点来作为起始节点
	idx := atomic.AddUint64(&f.idx, 1)
	length := uint64(len(f.svcs))
	for i := idx; i < idx+length; i++ {
		svc := f.svcs[int(i%length)]
		err := svc.Send(ctx, tpl, args, numbers...)
		switch err {
		case nil:
			return nil
		//	超时和被主动取消了
		case context.DeadlineExceeded, context.Canceled:
			return err
		default:
			//	输出日志
		}
	}
	return errors.New("全部服务商都失败了")
}

func NewFailoverSMSService(svcs []sms.Service) sms.Service {
	return &FailoverSMSService{
		svcs: svcs,
	}
}
