package repository

import (
	"basic-go/webook/internal/domain"
	"basic-go/webook/internal/repository/dao"
	"context"
	"time"
)

var (
	ErrUserDuplicateEmail = dao.ErrUserDuplocateEmail
	ErrUserNotFound       = dao.ErrUserNotFound
)

//var ErrUserDuplicateEmailV1 = fmt.Errorf("%w 邮箱冲突",dao.ErrUserDuplocateEmail)

type UserRepository struct {
	dao *dao.UserDao
}

func NewUserRepository(dao *dao.UserDao) *UserRepository {
	return &UserRepository{
		dao: dao,
	}
}

func (r *UserRepository) FindByemail(ctx context.Context, email string) (domain.User, error) {
	u, err := r.dao.FindByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{
		Id:       u.Id,
		Email:    u.Email,
		Password: u.Password,
	}, nil
}

func (r *UserRepository) FindByid(ctx context.Context, id int) (domain.User, error) {
	u, err := r.dao.FindById(ctx, id)
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{
		Id:          u.Id,
		Email:       u.Email,
		Name:        u.Name,
		Birthday:    u.Birthday,
		Description: u.Description,
	}, nil
}

func (r *UserRepository) Create(ctx context.Context, u domain.User) error {
	return r.dao.Insert(ctx, dao.User{
		Email:    u.Email,
		Password: u.Password,
	})

	//	操作缓存的位置
}

func (r *UserRepository) Update(ctx context.Context, u domain.User) error {
	return r.dao.Update(ctx, dao.User{
		Id:          u.Id,
		Name:        u.Name,
		Birthday:    u.Birthday,
		Description: u.Description,
		Utime:       time.Now().UnixMilli(),
	})

}

func (r *UserRepository) FindById(int64) {
	//先从cache中找
	//再从dao中找
	//找到了回写cache
}
