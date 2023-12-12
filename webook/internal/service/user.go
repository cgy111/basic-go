package service

import (
	"basic-go/webook/internal/domain"
	"basic-go/webook/internal/repository"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var ErrUserDuplicateEmail = repository.ErrUserDuplicateEmail
var ErrInvalidUserOrPassword = errors.New("账号/邮箱或密码不对")

type UserService struct {
	repo  *repository.UserRepository
	redis *redis.Client
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (svc *UserService) Login(ctx context.Context, email, password string) (domain.User, error) {
	//	先找用户
	u, err := svc.repo.FindByemail(ctx, email)
	if err == repository.ErrUserNotFound {
		return domain.User{}, ErrInvalidUserOrPassword
	}
	if err != nil {
		return domain.User{}, err
	}
	//	比较密码
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		//DEBUG
		return domain.User{}, ErrInvalidUserOrPassword
	}
	return u, nil
}

func (svc *UserService) SignUp(ctx context.Context, u domain.User) error {
	//考虑加密放在哪里
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	//然后存起来
	err = svc.repo.Create(ctx, u)
	if err != nil {
		return err
	}
	//redis不知道怎么处理这个u
	val, err := json.Marshal(u)
	if err != nil {
		return err
	}
	//注意这边，要求 u 的 id 不为 0
	err = svc.redis.Set(ctx, fmt.Sprintf("user Info:%d", u.Id), val, time.Minute*30).Err()
	return err

}

func (svc *UserService) Edit(ctx context.Context, u domain.User) error {

	return svc.repo.Update(ctx, u)
}

//	func (svc *UserService) Profile(ctx context.Context, id int) (domain.User, error) {
//		return svc.repo.FindByid(ctx, id)
//	}
func (svc *UserService) Profile(ctx context.Context, id int64) (domain.User, error) {
	//第一个念头先从缓存取
	val, err := svc.redis.Get(ctx, fmt.Sprintf("user info:%d", id)).Result()
	if err != nil {
		return domain.User{}, err
	}
	var u domain.User
	err = json.Unmarshal([]byte(val), &u)
	if err == nil {
		return u, err
	}
	//接下来从数据库中查找
	return svc.repo.FindByid(ctx, id)
}
