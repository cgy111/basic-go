package wechat

import (
	"basic-go/webook/internal/domain"
	"context"
	"encoding/json"
	"fmt"
	uuid "github.com/lithammer/shortuuid/v4"
	"net/http"
	"net/url"
)

var redirectURI = url.PathEscape("https://sxcgy.cn/wechat/callback")

type Service interface {
	AuthURL(ctx context.Context) (string, error)
	VerifyCode(ctx context.Context, code string, state string) (domain.WechatInfo, error)
}

type service struct {
	appId     string
	appSecret string
	cliect    *http.Client
}

// 不偷懒的写法
func NewServiceV1(appId string, appSecret string, client *http.Client) Service {
	return &service{
		appId:     appId,
		appSecret: appSecret,
		cliect:    client,
	}
}

func NewService(appId string, appSecret string) Service {
	return &service{
		appId:     appId,
		appSecret: appSecret,
		//依赖注入,没有完全注入
		cliect: http.DefaultClient,
	}
}

func (s *service) VerifyCode(ctx context.Context, code string, state string) (domain.WechatInfo, error) {
	//方法1
	//const targetPattern = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code\n"
	//taget := fmt.Sprintf(targetPattern, s.appId, s.appSecret, code)
	//resp, err := http.Get(taget)
	//if err != nil {
	//	return domain.WechatInfo{},nil
	//}
	//方法2
	const targetPattern = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code\n"
	taget := fmt.Sprintf(targetPattern, s.appId, s.appSecret, code)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, taget, nil)
	//req, err := http.NewRequest(http.MethodGet, taget, nil)
	if err != nil {
		return domain.WechatInfo{}, err
	}
	//会产生复制,性能极差,比如URL很长
	//req = req.WithContext(ctx)
	resp, err := s.cliect.Do(req)
	if err != nil {
		return domain.WechatInfo{}, err
	}

	//只读一遍
	decoder := json.NewDecoder(resp.Body)
	var res Result
	err = decoder.Decode(&res)

	//整个响应都读出来,不推荐，因为Unmarshal 再读一遍，合计两遍
	//body, err := io.ReadAll(resp.Body)
	//err = json.Unmarshal(body, &res)

	if err != nil {
		return domain.WechatInfo{}, err
	}
	if res.ErrCode != 0 {
		return domain.WechatInfo{}, fmt.Errorf("微信返回错误响应,错误码:%d,错误信息:%s", res.ErrCode, res.ErrMsg)
	}
	return domain.WechatInfo{
		OpenID:  res.Openid,
		UnionID: res.Unionid,
	}, nil
}

func (s *service) AuthURL(ctx context.Context) (string, error) {
	const urlPattern = "https://open.weixin.qq.com/connect/qrconnect?appid=%s&redirect_uri=%s&response_type=code&scope=snsapi_login&state=%s#wechat_redirect"
	state := uuid.New()
	return fmt.Sprintf(urlPattern, s.appId, redirectURI, state), nil
}

type Result struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`

	Assesstoken  string `json:"access_token"`
	Expiresin    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`

	Openid  string `json:"openid"`
	Scope   string `json:"scope"`
	Unionid string `json:"unionid"`
}
