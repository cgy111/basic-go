package dao

import (
	"context"
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	gormMysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestGORMUserDao_Insert(t *testing.T) {
	testCase := []struct {
		name string

		//为什么不用ctrl ？
		//因为这里是sqlmock，不是gomock
		mock    func(t *testing.T) *sql.DB
		ctx     context.Context
		user    User
		wantErr error
		//wantId  int64
	}{
		{
			name: "插入成功",
			mock: func(t *testing.T) *sql.DB {
				mockDB, mock, err := sqlmock.New()
				res := sqlmock.NewResult(3, 1)
				//这边输入的是正则表达式
				//这个写法的意思是，只要是 INSERT 到 users的语句
				mock.ExpectExec("INSERT INTO `users`.*").WillReturnResult(res)
				require.NoError(t, err)
				return mockDB
			},
			user: User{
				Email: sql.NullString{
					String: "123@qq.com",
					Valid:  true,
				},
			},
			//wantId: 3,
		},

		{
			name: "邮箱冲突",
			mock: func(t *testing.T) *sql.DB {
				mockDB, mock, err := sqlmock.New()
				//这边输入的是正则表达式
				//这个写法的意思是，只要是 INSERT 到 users的语句
				mock.ExpectExec("INSERT INTO `users`.*").WillReturnError(&mysql.MySQLError{
					Number: 1062,
				})
				require.NoError(t, err)
				return mockDB
			},
			user: User{},
			//wantId: 3,
			wantErr: ErrUserDuplocate,
		},

		{
			name: "数据库冲突",
			mock: func(t *testing.T) *sql.DB {
				mockDB, mock, err := sqlmock.New()
				//这边输入的是正则表达式
				//这个写法的意思是，只要是 INSERT 到 users的语句
				mock.ExpectExec("INSERT INTO `users`.*").WillReturnError(errors.New("数据库冲突"))
				require.NoError(t, err)
				return mockDB
			},
			user: User{},
			//wantId: 3,
			wantErr: errors.New("数据库冲突"),
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {

			db, err := gorm.Open(gormMysql.New(gormMysql.Config{
				Conn: tc.mock(t),
				//select version;
				SkipInitializeWithVersion: true,
			}), &gorm.Config{
				//mock DB 不需要ping
				DisableAutomaticPing:   true,
				SkipDefaultTransaction: true,
			})
			d := NewUserDAO(db)
			err = d.Insert(tc.ctx, tc.user)
			assert.Equal(t, tc.wantErr, err)
			//可以比较一下
			//assert.Equal(t, tc.wantId, tc.user.Id)
		})

	}
}
