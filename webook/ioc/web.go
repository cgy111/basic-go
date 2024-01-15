package ioc

import (
	"basic-go/webook/internal/web"
	"basic-go/webook/internal/web/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func InitGin(mdls []gin.HandlerFunc, hdl *web.UserHandler,
	oauth2WechatHdl *web.OAuth2WechatHandler) *gin.Engine {
	server := gin.Default()
	server.Use(mdls...)
	hdl.RegisterRoutes(server)
	oauth2WechatHdl.RegisterRoutes(server)
	return server
}

// redisClient redis.Cmdable
func InitMiddlewares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		// 跨域处理
		corsHdl(),

		// 登录验证
		middleware.NewLoginJWTMiddlewareBuilder().
			IgnorePaths("/users/signup").
			IgnorePaths("/users/login_sms/code/send").
			IgnorePaths("/users/login_sms").
			IgnorePaths("/oauth2/wechat/authurl").
			IgnorePaths("/oauth2/wechat/callback").
			IgnorePaths("/users/login").Build(),

		// 限流处理
		//ratelimit.NewBuilder(redisClient,time.Second,100).Build(),
	}
}

func corsHdl() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"POST", "GET"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
		//不加，前端拿不到
		ExposeHeaders: []string{"x-jwt-token"},
		//是否允许你带cookie之类的东西
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost:") {
				//开发环境
				return true
			}
			return strings.Contains(origin, "qiniu.io")
		},
		MaxAge: 12 * time.Hour,
	})
}
