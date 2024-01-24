// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"basic-go/webook/internal/repository"
	"basic-go/webook/internal/repository/cache"
	"basic-go/webook/internal/repository/dao"
	"basic-go/webook/internal/service"
	"basic-go/webook/internal/web"
	"basic-go/webook/ioc"
	"github.com/gin-gonic/gin"
)

import (
	_ "github.com/gin-contrib/sessions/redis"
)

// Injectors from wire.go:

func InitWebServer() *gin.Engine {
	v := ioc.InitMiddlewares()
	db := ioc.InitDB()
	userDAO := dao.NewUserDAO(db)
	cmdable := ioc.InitRedis()
	userCache := cache.NewUserCache(cmdable)
	userRepository := repository.NewUserRepository(userDAO, userCache)
	userService := service.NewUserService(userRepository)
	codeCache := cache.NewCodeCache(cmdable)
	codeRepository := repository.NewCodeRepository(codeCache)
	smsService := ioc.InitSmsService()
	codeService := service.NewCodeService(codeRepository, smsService)
	userHandler := web.NewUserHandler(userService, codeService)
	wechatService := ioc.InitWechatService()
	oAuth2WechatHandler := web.NewOAuth2WechatHandler(wechatService, userService)
	engine := ioc.InitGin(v, userHandler, oAuth2WechatHandler)
	return engine
}
