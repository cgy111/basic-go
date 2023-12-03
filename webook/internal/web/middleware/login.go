package middleware

import (
	"encoding/gob"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type LoginMiddlewareBuilder struct {
	paths []string
}

func NewLoginMiddlewareBuilder() *LoginMiddlewareBuilder {
	return &LoginMiddlewareBuilder{}
}

func (l *LoginMiddlewareBuilder) IgnorePaths(path string) *LoginMiddlewareBuilder {
	l.paths = append(l.paths, path)
	return l
}

func (l *LoginMiddlewareBuilder) Build() gin.HandlerFunc {
	//用Go的方式编码解码
	gob.Register(time.Now())
	return func(ctx *gin.Context) {
		for _, path := range l.paths {
			if ctx.Request.URL.Path == path {
				return
			}
		}

		//这些不需要登录校验
		//if ctx.Request.URL.Path == "/users/login" ||
		//	ctx.Request.URL.Path == "/users/signup" {
		//	return
		//}

		sess := sessions.Default(ctx)
		/*if sess != nil {
		//	没有登录
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}*/

		id := sess.Get("userId")

		if id == nil {
			//	没有登录
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		//拿到上一次更新时间
		updateTime := sess.Get("update_time")
		sess.Set("userId", id)
		sess.Options(sessions.Options{
			MaxAge: 60,
		})
		now := time.Now()
		//说明还没有刷新过，刚登陆，还没有刷新过
		if updateTime != nil {
			sess.Set("update_time", now)
			if err := sess.Save(); err != nil {
				panic(err)
			}

		}

		//updateTime是有的
		updateTimeVal, _ := updateTime.(time.Time)
		//if !ok {
		//	ctx.AbortWithStatus(http.StatusInternalServerError)
		//	return
		//}

		if now.Sub(updateTimeVal) > time.Second*10 {
			sess.Set("update_time", now)
			if err := sess.Save(); err != nil {
				panic(err)
			}
		}

		//fmt.Println("1")
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
