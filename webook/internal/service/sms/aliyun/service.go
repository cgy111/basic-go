package aliyun

import (
	"basic-go/webook/internal/service/sms"
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"strconv"
	"strings"
)

type Service struct {
	signName        string
	accessKeyId     string
	accessKeySecret string
	client          *dysmsapi.Client
}

func NewService(signName, accessKeyId, accessKeySecret string, client *dysmsapi.Client) *Service {
	return &Service{
		signName:        signName,
		accessKeyId:     accessKeyId,
		accessKeySecret: accessKeySecret,
		client:          client,
	}
}

func (s *Service) SendAliyun(tpl string, args []string, numbers ...string) error {
	req := dysmsapi.CreateSendSmsRequest()
	req.Scheme = "https"
	req.PhoneNumbers = strings.Join(numbers, ",")
	req.SignName = s.signName
	req.TemplateCode = tpl
	//传入JSON
	argsMap := make(map[string]string)
	for idx, arg := range args {
		argsMap[strconv.Itoa(idx)] = arg
	}
	//这就意味着短信验证码是{0}
	bCode, err := json.Marshal(argsMap)
	if err != nil {
		return err
	}
	req.TemplateParam = string(bCode)
	resp, err := s.client.SendSms(req)
	if err != nil {
		return err
	}
	if resp.Code != "OK" {
		return fmt.Errorf("发送短信失败: %s,%s", resp.Code, resp.Message)
	}
	return nil
}
func (s *Service) SendAliyunV1(tpl string, args []sms.NameArg, numbers ...string) error {
	req := dysmsapi.CreateSendSmsRequest()
	req.Scheme = "https"
	req.PhoneNumbers = strings.Join(numbers, ",")
	req.SignName = s.signName
	req.TemplateCode = tpl
	//传入JSON
	argsMap := make(map[string]string)
	for _, arg := range args {
		argsMap[arg.Name] = arg.Val
	}
	//这就意味着短信验证码是{0}
	bCode, err := json.Marshal(argsMap)
	if err != nil {
		return err
	}
	req.TemplateParam = string(bCode)
	resp, err := s.client.SendSms(req)
	if err != nil {
		return err
	}
	if resp.Code != "OK" {
		return fmt.Errorf("发送短信失败: %s,%s", resp.Code, resp.Message)
	}
	return nil
}

func (s *Service) SendAliyunV2(tpl string, args any, numbers ...string) error {
	req := dysmsapi.CreateSendSmsRequest()
	req.Scheme = "https"
	req.PhoneNumbers = strings.Join(numbers, ",")
	req.SignName = s.signName
	req.TemplateCode = tpl
	//传入JSON
	//这就意味着短信验证码是{0}
	bCode, err := json.Marshal(args.(map[string]string))
	if err != nil {
		return err
	}
	req.TemplateParam = string(bCode)
	resp, err := s.client.SendSms(req)
	if err != nil {
		return err
	}
	if resp.Code != "OK" {
		return fmt.Errorf("发送短信失败: %s,%s", resp.Code, resp.Message)
	}
	return nil
}
