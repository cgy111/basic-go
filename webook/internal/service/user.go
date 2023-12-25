package service

import (
	"basic-go/webook/internal/domain"
	"basic-go/webook/internal/repository"
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
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
	u, err := svc.repo.FindByEmail(ctx, email)
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
	return svc.repo.Create(ctx, u)
}

func (svc *UserService) Edit(ctx context.Context, u domain.User) error {

	return svc.repo.Update(ctx, u)
}

//	func (svc *UserService) Profile(ctx context.Context, id int) (domain.User, error) {
//		return svc.repo.FindByid(ctx, id)
//	}

func (svc *UserService) FindOrCreate(ctx context.Context, phone string) (domain.User, error) {
	//panic("先不实现")
	u, err := svc.repo.FindByPhone(ctx, phone)
	//要判断，有没有这个用户
	if err != repository.ErrUserNotFound {
		//nil会进了这里
		//不为ErrUserNotFound的也会进来这里
		return u, err
	}
	//明确知道没有这个用户，创建用户
	u = domain.User{
		Phone: phone,
	}
	err = svc.repo.Create(ctx, u)
	if err != nil {
		return u, err
	}
	//因为这里会遇到主从延迟问题
	return svc.repo.FindByPhone(ctx, phone)
}

func (svc *UserService) Profile(ctx context.Context, id int64) (domain.User, error) {
	u, err := svc.repo.FindById(ctx, id)
	return u, err
}
