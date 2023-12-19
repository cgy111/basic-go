package service

import "context"

type CodeService struct {
}

// Send 发送验证码
func (svc *CodeService) Send(ctx context.Context,
	//区别业务场景
	biz string,
	phone string) error {

}

// Verify 验证验证码
func (svc *CodeService) Verify(ctx context.Context, biz string,
	phone string, code string) (bool, error) {

}

func (svc *CodeService) VerifyV1(ctx context.Context, biz string,
	phone string, code string) error {

}
