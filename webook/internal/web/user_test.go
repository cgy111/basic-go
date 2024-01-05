package web

import (
	"basic-go/webook/internal/domain"
	"basic-go/webook/internal/service"
	svcmocks "basic-go/webook/internal/service/mocks"
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEncrypt(t *testing.T) {
	password := "hello#world123"
	encrypted, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		t.Fatal(err)
	}

	err = bcrypt.CompareHashAndPassword(encrypted, []byte(password))
	assert.NoError(t, err)
}

func TestNil(t *testing.T) {
	testTypeAssert(nil)

}

func testTypeAssert(c any) {
	claims := c.(*UserClaims)
	println(claims.Uid)
}

func TestUserHandler_Signup(t *testing.T) {
	testCases := []struct {
		name string

		mock func(ctrl *gomock.Controller) service.UserService

		reqBody string

		wantCode int

		wantBody string
	}{
		{
			name: "注册成功",
			mock: func(ctrl *gomock.Controller) service.UserService {
				usersvc := svcmocks.NewMockUserService(ctrl)
				usersvc.EXPECT().SignUp(gomock.Any(), domain.User{
					Email:    "123cgy@qq.com",
					Password: "hello@world123",
				}).Return(nil)
				//注册成功return nil
				return usersvc
			},
			reqBody: `
{
"email":"123cgy@qq.com",
"password":"hello@world123",
"ConfirmPassword":"hello@world123"
}
`,
			wantCode: http.StatusOK,
			wantBody: "注册成功",
		},

		{
			name: "参数不对，bind失败",
			mock: func(ctrl *gomock.Controller) service.UserService {
				usersvc := svcmocks.NewMockUserService(ctrl)
				//注册成功return nil
				return usersvc
			},
			reqBody: `
{
"email":"123hjx@qq.com",
"password":"hello@world123",
}
`,
			wantCode: http.StatusBadRequest,
		},

		{
			name: "邮箱格式不对",
			mock: func(ctrl *gomock.Controller) service.UserService {
				usersvc := svcmocks.NewMockUserService(ctrl)
				return usersvc
			},

			reqBody: `
		{
			"email": "123@q",
			"password": "hello#world123",
			"confirmPassword": "hello#world123"
		}
		`,
			wantCode: http.StatusOK,
			wantBody: "你的邮箱格式不对",
		},

		{
			name: "两次输入密码不匹配",
			mock: func(ctrl *gomock.Controller) service.UserService {
				usersvc := svcmocks.NewMockUserService(ctrl)
				//注册成功return nil
				return usersvc
			},
			reqBody: `
{
"email":"123cgy@qq.com",
"password":"hello@world1234",
"ConfirmPassword":"hello@world123"
}
`,
			wantCode: http.StatusOK,
			wantBody: "两次输入的密码不一致",
		},

		{
			name: "密码格式不对",
			mock: func(ctrl *gomock.Controller) service.UserService {
				usersvc := svcmocks.NewMockUserService(ctrl)
				//注册成功return nil
				return usersvc
			},
			reqBody: `
{
"email":"123cgy@qq.com",
"password":"hello",
"ConfirmPassword":"hello"
}
`,
			wantCode: http.StatusOK,
			wantBody: "密码必须大于8位，包含数字、特殊字符",
		},

		{
			name: "邮箱冲突",
			mock: func(ctrl *gomock.Controller) service.UserService {
				usersvc := svcmocks.NewMockUserService(ctrl)
				usersvc.EXPECT().SignUp(gomock.Any(), domain.User{
					Email:    "123cgy@qq.com",
					Password: "hello@world123",
				}).Return(service.ErrUserDuplicate)
				//注册成功return nil
				return usersvc
			},
			reqBody: `
{
"email":"123cgy@qq.com",
"password":"hello@world123",
"ConfirmPassword":"hello@world123"
}
`,
			wantCode: http.StatusOK,
			wantBody: "邮箱冲突",
		},

		{
			name: "系统异常",
			mock: func(ctrl *gomock.Controller) service.UserService {
				usersvc := svcmocks.NewMockUserService(ctrl)
				usersvc.EXPECT().SignUp(gomock.Any(), domain.User{
					Email:    "123cgy@qq.com",
					Password: "hello@world123",
				}).Return(errors.New("未知异常"))
				//注册成功return nil
				return usersvc
			},
			reqBody: `
{
"email":"123cgy@qq.com",
"password":"hello@world123",
"ConfirmPassword":"hello@world123"
}
`,
			wantCode: http.StatusOK,
			wantBody: "系统异常",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			//	这里怎么拿到这个响应
			/*		handler := NewUserHandler(nil, nil)
					ctx := &gin.Context{}
					handler.Signup(ctx)*/

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			server := gin.Default()
			//用不上 codeSvc
			h := NewUserHandler(tc.mock(ctrl), nil)
			h.RegisterRoutes(server)

			req, err := http.NewRequest(http.MethodPost,
				"/users/signup", bytes.NewBuffer([]byte(tc.reqBody)))
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
			assert.Equal(t, tc.wantBody, resp.Body.String())
		})
	}
}

//func TestMock(t * testing.T){
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	usersvc := svcmocks.NewMockUserService(ctrl)
//
//	usersvc.EXPECT().SignUp(gomock.Any(), gomock.Any()).
//		Return(errors.New("mock error"))
//
//	//usersvc.EXPECT().SignUp(gomock.Any(), domain.User{
//	//	Email: "123@qq.com",
//	//}).Return(errors.New("mock error"))
//
//	err := usersvc.SignUp(context.Background(), domain.User{
//		Email: "1234@qq.com",
//	})
//	t.Log(err)
//}
