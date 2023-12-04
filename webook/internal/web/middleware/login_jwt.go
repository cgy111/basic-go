package middleware

import (
	"encoding/gob"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// JWT登录校验
type LoginJWTMiddlewareBuilder struct {
	paths []string
}

func NewLoginJWTMiddlewareBuilder() *LoginMiddlewareBuilder {
	return &LoginMiddlewareBuilder{}
}

/*func (l *LoginJWTMiddlewareBuilder) IgnorePaths(path string) *LoginJWTMiddlewareBuilder {
	l.paths = append(l.paths, path)
	return l
}*/

func (l *LoginJWTMiddlewareBuilder) Build() gin.HandlerFunc {
	//用Go的方式编码解码
	gob.Register(time.Now())
	return func(ctx *gin.Context) {
		for _, path := range l.paths {
			if ctx.Request.URL.Path == path {
				return
			}
		}
		//现在用JWT来校验
		token := ctx.GetHeader("Authorization")
		if token == "" {
			//	没有登录
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

/*func (l *LoginMiddlewareBuilder) CheckLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//这些不需要登录校验
		if ctx.Request.URL.Path == "/users/login" ||
			ctx.Request.URL.Path == "/users/signup" {
			return
		}

		sess := sessions.Default(ctx)
		///*if sess != nil {
		////	没有登录
		//	ctx.AbortWithStatus(http.StatusUnauthorized)
		//	return
		//}
		id := sess.Get("userId")
		if id == nil {
			//	没有登录
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		fmt.Println("1")
	}
}*/
