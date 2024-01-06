package repository

import (
	"basic-go/webook/internal/domain"
	"basic-go/webook/internal/repository/cache"
	cachemocks "basic-go/webook/internal/repository/cache/mocks"
	"basic-go/webook/internal/repository/dao"
	daomocks "basic-go/webook/internal/repository/dao/mocks"
	"context"
	"database/sql"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
	"time"
)

func TestCachedUserRepository_FindById(t *testing.T) {
	now := time.Now()
	now = time.UnixMilli(now.UnixMilli())
	testCases := []struct {
		name string

		mock func(ctrl *gomock.Controller) (dao.UserDAO, cache.UserCache)

		ctx context.Context
		id  int64

		wantUser domain.User
		wantErr  error
	}{
		{
			name: "缓存未命中,但是查询成功",
			mock: func(ctrl *gomock.Controller) (dao.UserDAO, cache.UserCache) {
				//缓存未命中，查了缓存，但是没结果
				c := cachemocks.NewMockUserCache(ctrl)
				c.EXPECT().Get(gomock.Any(), int64(123)).
					Return(domain.User{}, cache.ErrNotExist)
				d := daomocks.NewMockUserDAO(ctrl)
				d.EXPECT().FindById(gomock.Any(), int64(123)).
					Return(dao.User{
						Id: 123,
						Email: sql.NullString{
							String: "123@qq.com",
							Valid:  true,
						},
						Password: "123",
						Phone: sql.NullString{
							String: "12345678901",
							Valid:  true,
						},
						Ctime: now.UnixMilli(),
						Utime: now.UnixMilli(),
					}, nil)

				c.EXPECT().Set(gomock.Any(), domain.User{
					Id:       123,
					Email:    "123@qq.com",
					Password: "123",
					Phone:    "12345678901",
					Ctime:    now,
				}).Return(nil)
				return d, c
			},
			ctx: context.Background(),
			id:  123,
			wantUser: domain.User{
				Id:       123,
				Email:    "123@qq.com",
				Password: "123",
				Phone:    "12345678901",
				Ctime:    now,
			},
			wantErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			ud, uc := tc.mock(ctrl)
			repo := NewUserRepository(ud, uc)
			u, err := repo.FindById(tc.ctx, tc.id)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.wantUser, u)
		})
	}
}
