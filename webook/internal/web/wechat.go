package web

import "github.com/gin-gonic/gin"

type OAuth2WechatHandler struct {
}

func NewOAuth2WechatHandler() {

}

func (h *OAuth2WechatHandler) RegisterRoutes(server *gin.Engine) {
	g := server.Group("/oauth2/wechat")
	g.GET("/authurl", h.AuthURL)
	g.Any("/callback", h.Callback)
}

func (h *OAuth2WechatHandler) AuthURL(ctx *gin.Context) {

}

func (h *OAuth2WechatHandler) Callback(ctx *gin.Context) {

}

type OAuth2Handler struct {
}

func (h *OAuth2Handler) RegisterRoutes(server *gin.Engine) {
	g := server.Group("/oauth2")
	g.GET("/:platform/authurl", h.AuthURL)
	g.Any("/:platform/callback", h.Callback)
}

func (h *OAuth2Handler) AuthURL(ctx *gin.Context) {
	platform := ctx.Param("platform")
	// 处理平台参数
	switch platform {
	case "wechat":

	}
}

func (h *OAuth2Handler) Callback(ctx *gin.Context) {

}
