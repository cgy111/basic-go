package integration

import (
	"basic-go/webook/internal/web"
	"basic-go/webook/ioc"
	"bytes"
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestUserHandler_e2e_SendLoginSmsCode(t *testing.T) {
	server := InitWebServer()
	rdb := ioc.InitRedis()
	testCases := []struct {
		name string

		//要考虑准备数据
		before func(t *testing.T)

		//以及验证数据,数据库的数据对不对,redis的数据对不对
		after    func(t *testing.T)
		reqBody  string
		wantCode int
		wantBody web.Result
	}{
		{
			name: "发送成功",
			before: func(t *testing.T) {
				//不需要，也就是Redis什么数据都没有
			},
			after: func(t *testing.T) {
				ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
				//清理数据
				val, err := rdb.GetDel(ctx, "phone_code:login:13800138000").Result()
				cancel()
				assert.NoError(t, err)
				//验证码是6位
				assert.True(t, len(val) == 6)
			},
			reqBody: `
{
    "phone": "13800138000"
}
`,
			wantCode: 200,
			wantBody: web.Result{
				Msg: "发送成功",
			},
		},

		{
			name: "发送太频繁",
			before: func(t *testing.T) {
				//这个手机号，已经有一个验证码了
				ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
				_, err := rdb.Set(ctx, "phone_code:login:13800138000", "123456",
					9*time.Minute+30*time.Second).Result()
				cancel()
				require.NoError(t, err)
			},
			after: func(t *testing.T) {
				ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
				//清理数据
				val, err := rdb.GetDel(ctx, "phone_code:login:13800138000").Result()
				cancel()
				assert.NoError(t, err)
				//验证码是6位,没有被覆盖,还是123456
				assert.Equal(t, "123456", val)
			},
			reqBody: `
{
    "phone": "13800138000"
}
`,
			wantCode: 200,
			wantBody: web.Result{
				Msg: "发送太频繁，请稍后再试",
			},
		},

		{
			name: "系统错误",
			before: func(t *testing.T) {
				//这个手机号，已经有一个验证码了
				ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
				_, err := rdb.Set(ctx, "phone_code:login:13800138000", "123456", 0).Result()
				cancel()
				require.NoError(t, err)
			},
			after: func(t *testing.T) {
				ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
				//清理数据
				val, err := rdb.GetDel(ctx, "phone_code:login:13800138000").Result()
				cancel()
				assert.NoError(t, err)
				//验证码是6位,没有被覆盖,还是123456
				assert.Equal(t, "123456", val)
			},
			reqBody: `
{
    "phone": "13800138000"
}
`,
			wantCode: 200,
			wantBody: web.Result{
				Code: 5,
				Msg:  "系统错误",
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(t)
			req, err := http.NewRequest(http.MethodPost,
				"/users/login_sms/code/send", bytes.NewBuffer([]byte(tc.reqBody)))
			require.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")
			//可以继续使用req
			resp := httptest.NewRecorder()
			//resp.Header()
			t.Log(resp)

			//这就是http请求进入GIN框架的入口
			//当你这样调用的时候，GIN就会处理这个请求
			//响应写回到resp里
			server.ServeHTTP(resp, req)

			assert.Equal(t, tc.wantCode, resp.Code)
			var webRes web.Result
			err = json.NewDecoder(resp.Body).Decode(&webRes)
			require.NoError(t, err) //预期不会出错
			assert.Equal(t, tc.wantBody, webRes)
			tc.after(t)
		})
	}
}
