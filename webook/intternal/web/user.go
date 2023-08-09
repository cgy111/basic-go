package web

import "github.com/gin-gonic/gin"

// UserHandler 定义所有和用户有关的路由
type UserHandler struct {
}

func (u *UserHandler) RegisterRoutes(server *gin.Engine) {
	//u := &UserHandler{}

	server.POST("/users/signup", u.SignUp)
	server.POST("/users/login", u.Login)
	//非Rest风格
	server.POST("/users/edit", u.Edit)
	server.GET("users/profile", u.Profile)

}

type ArticleHandler struct {
}

func (u *UserHandler) SignUp(ctx *gin.Context) {

}

func (u *UserHandler) Login(ctx *gin.Context) {

}

func (u *UserHandler) Edit(ctx *gin.Context) {

}

func (u *UserHandler) Profile(ctx *gin.Context) {

}
