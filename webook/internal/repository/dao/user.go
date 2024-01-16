package dao

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	ErrUserDuplocate = errors.New("邮箱冲突或手机号冲突")
	ErrUserNotFound  = gorm.ErrRecordNotFound
)

type UserDAO interface {
	FindByEmail(ctx context.Context, email string) (User, error)
	FindByPhone(ctx context.Context, phone string) (User, error)
	FindById(ctx context.Context, id int64) (User, error)
	Insert(ctx context.Context, u User) error
	Update(ctx context.Context, u User) error
	FindByWechat(ctx context.Context, openId string) (User, error)
}

type GORMUserDao struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) UserDAO {
	return &GORMUserDao{
		db: db,
	}
}

func (dao *GORMUserDao) FindByWechat(ctx context.Context, openId string) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).Where("wechat_open_id=?", openId).First(&u).Error
	//err := dao.db.WithContext(ctx).First(&u,"email=?",email).Error
	return u, err
}

func (dao *GORMUserDao) FindByEmail(ctx context.Context, email string) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).Where("email=?", email).First(&u).Error
	//err := dao.db.WithContext(ctx).First(&u,"email=?",email).Error
	return u, err
}

func (dao *GORMUserDao) FindByPhone(ctx context.Context, phone string) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).Where("phone=?", phone).First(&u).Error
	//err := dao.db.WithContext(ctx).First(&u,"email=?",email).Error
	return u, err
}

func (dao *GORMUserDao) FindById(ctx context.Context, id int64) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).Where("'id' = ?", id).First(&u).Error
	return u, err

}

func (dao *GORMUserDao) Insert(ctx context.Context, u User) error {
	//存毫秒数
	now := time.Now().UnixMilli()
	u.Utime = now
	u.Ctime = now
	err := dao.db.WithContext(ctx).Create(&u).Error
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		const uniqueConfictsErrNo uint16 = 1062
		if mysqlErr.Number == uniqueConfictsErrNo {
			//邮箱冲突或手机号码冲突
			return ErrUserDuplocate
		}
	}

	return err
}

func (dao *GORMUserDao) Update(ctx context.Context, u User) error {
	err := dao.db.WithContext(ctx).Updates(&u).Error
	return err
}

// User直接对应数据库表结构
// 有的叫entity，model,PO(persistent object)
type User struct {
	Id int64 `gorm:"primaryKey,autoIncrement"`
	//全部用户唯一
	Email       sql.NullString `gorm:"unique"`
	Password    string
	Name        string
	Phone       sql.NullString `gorm:"unique"`
	Birthday    string
	Description string

	//如果要创建联合索引，<unionid,openid>,用openid查询的时候不会走索引
	//<openid,unionid>用 unionid 查询的时候，不会走索引

	//微信的字段
	WechatUnionID sql.NullString
	WechatOpenID  sql.NullString `gorm:"unique"`

	//创建时间，毫秒数
	Ctime int64
	//更新时间，毫秒数
	Utime int64
}
