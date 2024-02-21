package web

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"net/http"
	"strings"
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

func (h jwtHandler) setLoginToken(ctx *gin.Context, uid int64) error {
	ssid := uuid.New().String()
	err := h.setJWTToken(ctx, uid, ssid)
	if err != nil {
		return err
	}
	err = h.setRefreshJWTToken(ctx, uid, ssid)
	return err
}

func (h jwtHandler) setJWTToken(ctx *gin.Context, uid int64, ssid string) error {

	claims := UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 30)),
		},
		Uid:       uid,
		Ssid:      ssid,
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

func (h jwtHandler) setRefreshJWTToken(ctx *gin.Context, uid int64, ssid string) error {
	claims := RefreshClaims{
		Ssid: ssid,
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

func ExtractToken(ctx *gin.Context) string {
	//现在用JWT来校验
	tokenHeader := ctx.GetHeader("Authorization")
	//segs := strings.SplitN(tokenHeader, " ", 2)
	segs := strings.Split(tokenHeader, " ")
	if len(segs) != 2 {
		//	没登录，有人瞎搞
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return ""
	}
	return segs[1]
}

type RefreshClaims struct {
	Uid  int64
	Ssid string
	jwt.RegisteredClaims
}

type UserClaims struct {
	jwt.RegisteredClaims
	//声明要放进token里面的数据
	Uid       int64
	Ssid      string
	UserAgent string
}