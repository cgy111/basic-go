package web

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type jwtHandler struct {
	//access_token key
	atKey []byte
	//refresh_token key
	rtKey []byte
}

func newJwtHandler() jwtHandler {
	return jwtHandler{
		atKey: []byte("8b8d2e454737a253e0b12365a1ab97e1"),
		rtKey: []byte("8b8d2e454737a253e0b12365a1ab97e3"),
	}
}

func (h jwtHandler) setJWTToken(ctx *gin.Context, uid int64) error {
	claims := UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 30)),
		},
		Uid:       uid,
		UserAgent: ctx.Request.UserAgent(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenStr, err := token.SignedString(h.atKey)
	if err != nil {
		//ctx.String(http.StatusInternalServerError, "系统错误")
		return err
	}
	ctx.Header("x-jwt-token", tokenStr)
	return nil
}

func (h jwtHandler) setRefreshJWTToken(ctx *gin.Context, uid int64) error {
	claims := RefreshClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 30)),
		},
		Uid: uid,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenStr, err := token.SignedString(h.rtKey)
	if err != nil {
		//ctx.String(http.StatusInternalServerError, "系统错误")
		return err
	}
	ctx.Header("x-refresh-token", tokenStr)
	return nil
}

type RefreshClaims struct {
	Uid int64
	jwt.RegisteredClaims
}

type UserClaims struct {
	jwt.RegisteredClaims
	//声明要放进token里面的数据
	Uid       int64
	UserAgent string
}
