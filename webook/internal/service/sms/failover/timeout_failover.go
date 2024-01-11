package failover

import (
	"basic-go/webook/internal/service/sms"
	"context"
	"sync/atomic"
)

type TimeoutFailoverSMSService struct {
	//服务商
	svcs []sms.Service
	idx  int32
	//连续超时个数
	cnt int32

	//阈值
	//连续超时个数大于阈值时，切换到下一个服务商
	threshold int32
}

func (t *TimeoutFailoverSMSService) Send(ctx context.Context, tpl string, args []string, numbers ...string) error {
	//idx := atomic.AddInt32(&t.idx, 1) % int32(len(t.svcs))
	idx := atomic.LoadInt32(&t.idx)
	cnt := atomic.LoadInt32(&t.cnt)
	if cnt < t.threshold {
		//这里要切换,这是新的下标，往后挪了一个
		newIdx := (idx + 1) % int32(len(t.svcs))
		if atomic.CompareAndSwapInt32(&t.idx, idx, newIdx) {
			//成功往后挪了一位
			atomic.StoreInt32(&t.cnt, 0)
		}
		//else 就是出现并发了，别人换成功了
		//切换到下一个服务商生效
		//idx = newIdx
		//不管切换几个都生效
		idx = atomic.LoadInt32(&t.idx)
	}
	svc := t.svcs[idx]
	err := svc.Send(ctx, tpl, args, numbers...)
	switch err {
	//超时
	case context.DeadlineExceeded:
		atomic.AddInt32(&t.cnt, 1)
		return err
	case nil:
		//连续状态被打断
		atomic.StoreInt32(&t.cnt, 0)
		return nil
	default:
		//未知错误

		//可以考虑换下一个
		//超时，可能是偶发，尽量再试试
		//非超时，直接下一个
		return err
	}
}

func NewTimeoutFailoverSMSService() sms.Service {
	return &TimeoutFailoverSMSService{}
}
