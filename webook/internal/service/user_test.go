package service

import (
	"basic-go/webook/internal/domain"
	"basic-go/webook/internal/repository"
	repomocks "basic-go/webook/internal/repository/mocks"
	"context"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"
	"testing"
	"time"
)

func Test_userService_Login(t *testing.T) {
	//做成一个测试用例都用到的时间
	now := time.Now()
	testCase := []struct {
		name string
		mock func(ctrl *gomock.Controller) repository.UserRepository
		//输入
		//ctx      context.Context
		email    string
		password string

		wantUser domain.User
		wantErr  error
	}{
		{
			name: "登录成功", //用户名和密码对得上
			mock: func(ctrl *gomock.Controller) repository.UserRepository {
				repo := repomocks.NewMockUserRepository(ctrl)
				repo.EXPECT().FindByEmail(gomock.Any(), "123cgy@qq.com").
					Return(domain.User{
						Email:    "123cgy@qq.com",
						Password: "$2a$10$ibW0Y0c5.YqZ.QaNIiZIieJxL8BDHAFriETJyVDTd.xf3SYBWba8e",
						Phone:    "12345678901",
						Ctime:    now,
					}, nil)
				return repo
			},
			email:    "123cgy@qq.com",
			password: "hello#world123",

			wantUser: domain.User{
				Email:    "123cgy@qq.com",
				Password: "$2a$10$ibW0Y0c5.YqZ.QaNIiZIieJxL8BDHAFriETJyVDTd.xf3SYBWba8e",
				Phone:    "12345678901",
				Ctime:    now,
			},
			wantErr: nil,
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			//具体测试代码
			svc := NewUserService(tc.mock(ctrl))
			u, err := svc.Login(context.Background(), tc.email, tc.password)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.wantUser, u)
		})
	}
}

func TestEncrypted(t *testing.T) {
	res, err := bcrypt.GenerateFromPassword([]byte("hello#world123"), bcrypt.DefaultCost)
	if err == nil {
		t.Log(string(res))
	}
	t.Log(bcrypt.GenerateFromPassword([]byte(res), bcrypt.DefaultCost))
}
