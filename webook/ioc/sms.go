package ioc

import (
	"basic-go/webook/internal/service/sms"
	"basic-go/webook/internal/service/sms/memory"
)

func InitSmsService() sms.Service {
	//可以随时替换
	return memory.NewService()
	//svc := memory.NewService()
	//return ratelimit.NewRatelimitSMSService(svc,)
}
