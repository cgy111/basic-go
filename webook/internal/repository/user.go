package repository

import (
	"basic-go/webook/internal/domain"
	"basic-go/webook/internal/repository/cache"
	"basic-go/webook/internal/repository/dao"
	"context"
	"database/sql"
	"time"
)

var (
	ErrUserDuplicate = dao.ErrUserDuplocate
	ErrUserNotFound  = dao.ErrUserNotFound
)

//var ErrUserDuplicateEmailV1 = fmt.Errorf("%w 邮箱冲突",dao.ErrUserDuplocateEmail)

type UserRepository interface {
	FindByPhone(ctx context.Context, phone string) (domain.User, error)
	FindByEmail(ctx context.Context, email string) (domain.User, error)
	FindById(ctx context.Context, id int64) (domain.User, error)
	Create(ctx context.Context, u domain.User) error
	Update(ctx context.Context, u domain.User) error
	FindByWechat(ctx context.Context, openID string) (domain.User, error)
}

type CachedUserRepository struct {
	dao   dao.UserDAO
	cache cache.UserCache
}

func NewUserRepository(dao dao.UserDAO, c cache.UserCache) UserRepository {
	return &CachedUserRepository{
		dao:   dao,
		cache: c,
	}
}
func (r *CachedUserRepository) FindByWechat(ctx context.Context, openID string) (domain.User, error) {
	u, err := r.dao.FindByWechat(ctx, openID)
	if err != nil {
		return domain.User{}, err
	}
	return r.entityToDomain(u), nil
}
func (r *CachedUserRepository) FindByPhone(ctx context.Context, phone string) (domain.User, error) {
	u, err := r.dao.FindByPhone(ctx, phone)
	if err != nil {
		return domain.User{}, err
	}
	return r.entityToDomain(u), nil
}

func (r *CachedUserRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	u, err := r.dao.FindByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return r.entityToDomain(u), nil
}

func (r *CachedUserRepository) FindById(ctx context.Context, id int64) (domain.User, error) {
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

	u = r.entityToDomain(ue)

	//_ = r.cache.Set(ctx, u)
	/*if err != nil {
		//这里要不要返回err
		//打日志，做监控
	}
	*/
	go func() {
		_ = r.cache.Set(ctx, u)
		/*if err != nil {
			//这里要不要返回err
			//打日志，做监控
		}*/
	}()
	return u, nil
	//err=io.EOF 要不要去数据库里面找？选加载，万一Redis真的崩掉了，你要保护住你的系统，
	//面试：选加载，数据库限流保护
	//实际：不加载，用户体验差一点

	//缓存里面有数据
	//缓存里面没有数据
	//缓存出错了，不知道有没有数据
}

func (r *CachedUserRepository) Create(ctx context.Context, u domain.User) error {
	return r.dao.Insert(ctx, r.domainToEntity(u))

	//	操作缓存的位置
}

func (r *CachedUserRepository) Update(ctx context.Context, u domain.User) error {
	return r.dao.Update(ctx, dao.User{
		Id:          u.Id,
		Name:        u.Name,
		Birthday:    u.Birthday,
		Description: u.Description,
		Utime:       time.Now().UnixMilli(),
	})

}

func (r *CachedUserRepository) domainToEntity(u domain.User) dao.User {
	return dao.User{
		Id: u.Id,
		Email: sql.NullString{
			String: u.Email,
			//确实有手机号
			Valid: u.Email != "",
		},
		Password: u.Password,
		Name:     u.Name,
		Phone: sql.NullString{
			String: u.Phone,
			Valid:  u.Phone != "",
		},
		Birthday:    u.Birthday,
		Description: u.Description,
		WechatOpenID: sql.NullString{
			String: u.WechatInfo.OpenID,
			Valid:  u.WechatInfo.OpenID != "",
		},
		WechatUnionID: sql.NullString{
			String: u.WechatInfo.UnionID,
			Valid:  u.WechatInfo.UnionID != "",
		},
		Ctime: u.Ctime.UnixMilli(),
	}
}

func (r *CachedUserRepository) entityToDomain(u dao.User) domain.User {
	return domain.User{
		Id:          u.Id,
		Email:       u.Email.String,
		Password:    u.Password,
		Name:        u.Name,
		Phone:       u.Phone.String,
		Birthday:    u.Birthday,
		Description: u.Description,

		// 索引的最左匹配原则：
		// 假如索引在 <A, B, C> 建好了
		// A, AB, ABC 都能用
		// WHERE A =?
		// WHERE A = ? AND B =?    WHERE B = ? AND A =?
		// WHERE A = ? AND B = ? AND C = ?  ABC 的顺序随便换
		// WHERE 里面带了 ABC，可以用
		// WHERE 里面，没有 A，就不能用

		// 如果要创建联合索引，<unionid, openid>，用 openid 查询的时候不会走索引
		// <openid, unionid> 用 unionid 查询的时候，不会走索引
		// 微信的字段

		WechatInfo: domain.WechatInfo{
			UnionID: u.WechatUnionID.String,
			OpenID:  u.WechatOpenID.String,
		},
		Ctime: time.UnixMilli(u.Ctime),
	}
}
