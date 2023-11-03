package service

import (
	"basic-go/webook/intternal/domain"
	"basic-go/webook/intternal/repository"
	"context"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (svc *UserService) SignUp(ctx context.Context, u domain.User) error {
	//考虑加密放在哪里
	//然后存起来
	return svc.repo.Create(ctx, u)
}
