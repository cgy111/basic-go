package service

import (
	"basic-go/webook/internal/repository"
	"basic-go/webook/internal/service/sms"
	"context"
	"fmt"
	"math/rand"
)

var (
	ErrorCodeSendTooMany      = repository.ErrorCodeSendTooMany
	ErrCodeVerifyTooManyTimes = repository.ErrCodeVerifyTooManyTimes
)

const codeTplId = "1877556"

type CodeService struct {
	repo   *repository.CodeRepository
	smsSvc sms.Service
}

func NewCodeService(repo *repository.CodeRepository, smsSvc sms.Service) *CodeService {
	return &CodeService{
		repo:   repo,
		smsSvc: smsSvc,
	}
}

// Send 发送验证码
func (svc *CodeService) Send(ctx context.Context,
	//区别业务场景
	biz string,
	phone string) error {
	//两个步骤，生成一个验证码
	code := svc.generateCode()
	//塞进Redis
	err := svc.repo.Store(ctx, biz, phone, code)
	if err != nil {
		return err
	}
	//发送出去
	err = svc.smsSvc.SendTencent(ctx, codeTplId, []string{code}, phone)
	return err
}

// Verify 验证验证码
func (svc *CodeService) Verify(ctx context.Context, biz string,
	phone string, inputCode string) (bool, error) {
	//phone_code:$biz:138xxxxxx
	//code:$biz:138xxxxxx
	//user:login:138xxxxxx
	return svc.repo.Verify(ctx, biz, phone, inputCode)
}

func (svc *CodeService) generateCode() string {
	//	六位数，num在0-99999之间，包含0和99999
	num := rand.Intn(1000000)
	//	不足六位，前面补0
	return fmt.Sprintf("%6d", num)
}

/*func (svc *CodeService) VerifyV1(ctx context.Context, biz string,
	phone string, code string) error {

}
*/
