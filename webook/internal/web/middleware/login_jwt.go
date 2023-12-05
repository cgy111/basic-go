package middleware

import (
	"basic-go/webook/internal/web"
	"encoding/gob"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"strings"
	"time"
)

// JWT登录校验
type LoginJWTMiddlewareBuilder struct {
	paths []string
}

func NewLoginJWTMiddlewareBuilder() *LoginJWTMiddlewareBuilder {
	return &LoginJWTMiddlewareBuilder{}
}

func (l *LoginJWTMiddlewareBuilder) IgnorePaths(path string) *LoginJWTMiddlewareBuilder {
	l.paths = append(l.paths, path)
	return l
}

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
		tokenHeader := ctx.GetHeader("Authorization")
		if tokenHeader == "" {
			//	没有登录
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		segs := strings.SplitN(tokenHeader, " ", 2)
		if len(segs) != 2 {
			//	没登录，有人瞎搞
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenStr := segs[1]
		claims := &web.UserClaims{}
		//ParseWithClaims里面，一定要传入指针
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("8b8d2e454737a253e0b12365a1ab97e2"), nil
		})

		if err != nil {
			//	没有登录
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		//claims.ExpiresAt.Time.Before(time.Now()){
		//	//过期了
		//}
		//err为nil，token不为nil
		if token == nil || !token.Valid || claims.Uid == 0 {
			//	没有登录
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute))
		tokenStr, err = token.SignedString([]byte("8b8d2e454737a253e0b12365a1ab97e2"))
		if err != nil {
			log.Print("jwt 续约失败", err)
		}
		ctx.Header("x-jwt-token", tokenStr)
		ctx.Set("claims", claims)
		//ctx.Set("userId",claims.Uid)
	}
}