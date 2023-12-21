package repository

import (
	"basic-go/webook/internal/repository/cache"
	"context"
)

var (
	ErrorCodeSendTooMany      = cache.ErrSendTooMany            //发送太频繁
	ErrCodeVerifyTooManyTimes = cache.ErrCodeVerifyTooManyTimes //验证太频繁
)

type CodeRepository struct {
	cache *cache.CodeCache
}

func NewCodeRepository(c *cache.CodeCache) *CodeRepository {
	return &CodeRepository{
		cache: c,
	}
}

func (repo *CodeRepository) Store(ctx context.Context, biz string,
	phone string, code string) error {
	return repo.cache.Set(ctx, biz, phone, code)
}

func (repo *CodeRepository) Verify(ctx context.Context, biz string,
	phone string, inputCode string) (bool, error) {
	return repo.cache.Verify(ctx, biz, phone, inputCode)
}