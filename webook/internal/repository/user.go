package repository

import (
	"basic-go/webook/internal/domain"
	"basic-go/webook/internal/repository/cache"
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
	dao   *dao.UserDao
	cache *cache.UserCache
}

func NewUserRepository(dao *dao.UserDao, c *cache.UserCache) *UserRepository {
	return &UserRepository{
		dao:   dao,
		cache: c,
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

//	func (r *UserRepository) FindByid(ctx context.Context, id int64) (domain.User, error) {
//		u, err := r.dao.FindById(ctx, id)
//		if err != nil {
//			return domain.User{}, err
//		}
//		return domain.User{
//			Id:          u.Id,
//			Email:       u.Email,
//			Name:        u.Name,
//			Birthday:    u.Birthday,
//			Description: u.Description,
//		}, nil
//	}
func (r *UserRepository) FindById(ctx context.Context, id int64) (domain.User, error) {
	u, err := r.cache.Get(ctx, id)
	if err == nil {
		return u, nil
	}
	//if err == cache.ErrNotExist {
	//	//	缓存里面没有数据,去数据库里面找
	//}

	ue, err := r.dao.FindById(ctx, id)
	if err != nil {
		return domain.User{}, err
	}

	u = domain.User{
		Id:          ue.Id,
		Email:       ue.Email,
		Name:        ue.Name,
		Birthday:    ue.Birthday,
		Description: ue.Description,
	}
	go func() {
		err = r.cache.Set(ctx, u)
		if err != nil {
			//这里要不要返回err
			//打日志，做监控
		}
	}()
	return u, nil
	//err=io.EOF 要不要去数据库里面找？选加载，万一Redis真的崩掉了，你要保护住你的系统，
	//面试：选加载，数据库限流保护
	//实际：不加载，用户体验差一点

	//缓存里面有数据
	//缓存里面没有数据
	//缓存出错了，不知道有没有数据
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

/*func (r *UserRepository) FindById(int64) {
	//先从cache中找
	//再从dao中找
	//找到了回写cache
}*/
