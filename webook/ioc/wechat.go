package ioc

import (
	"basic-go/webook/internal/service/oauth2/wechat"
	"os"
)

func InitWechatService() wechat.Service {
	appId, ok := os.LookupEnv("WECHAT_APP_ID")
	if !ok {
		panic("没有找到环境变量 WECHAT_APP_ID")
	}
	//fmt.Println(appId + "\n")
	appKey, ok := os.LookupEnv("WECHAT_APP_SECRET")
	if !ok {
		panic("没有找到环境变量 WECHAT_APP_SECRET")
	}
	//fmt.Println(appKey + "\n")
	return wechat.NewService(appId, appKey)
}
