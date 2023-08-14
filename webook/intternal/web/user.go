package web

import (
	"fmt"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserHandler 定义所有和用户有关的路由
type UserHandler struct {
	emailExp    *regexp.Regexp
	passwordExp *regexp.Regexp
}

func NewUserHandler() *UserHandler {
	const (
		emailRegexPattern    = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
		passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@!%*#?&]{8,}$`
	)
	emailExp := regexp.MustCompile(emailRegexPattern, regexp.None)
	passwordExp := regexp.MustCompile(passwordRegexPattern, regexp.None)
	return &UserHandler{
		emailExp:    emailExp,
		passwordExp: passwordExp,
	}
}

func (u *UserHandler) RegisterRoutes(server *gin.Engine) {
	//u := &UserHandler{}

	/*	server.POST("/users/signup", u.SignUp)
		server.POST("/users/login", u.Login)
		//非Rest风格
		server.POST("/users/edit", u.Edit)
		server.GET("users/profile", u.Profile)*/

	ug := server.Group("/users")

	ug.GET("/profile", u.Profile)
	ug.POST("/signup", u.SignUp)
	ug.POST("/login", u.Login)
	//ug.POST("/login/page", u.Login)
	ug.POST("/edit", u.Edit)
}

/*func (u *UserHandler) RegisterRoutesV1(ug *gin.RouterGroup) {
	//u := &UserHandler{}

		server.POST("/users/signup", u.SignUp)
		server.POST("/users/login", u.Login)
		//非Rest风格
		server.POST("/users/edit", u.Edit)
		server.GET("users/profile", u.Profile)

	ug.GET("/profile", u.Profile)
	ug.POST("/signup", u.SignUp)
	ug.POST("/login", u.Login)
	ug.POST("/edit", u.Edit)
}
*/

type ArticleHandler struct {
}

func (u *UserHandler) SignUp(ctx *gin.Context) {
	type SignUpReq struct {
		Email           string `json:"email"`
		ConfirmPassword string `json:"confirmPassword"`
		Password        string `json:"password"`
	}
	//Bind方法会根据Content-Type来解析你的数据到req里面
	//解析错了，就会直接写回一个400的错误
	var req SignUpReq
	if err := ctx.Bind(&req); err != nil {
		return
	}

	//邮箱校验
	ok, err := u.emailExp.MatchString(req.Email)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	if !ok {
		ctx.String(http.StatusOK, "你的邮箱格式不对")
		return
	}

	//密码校验
	if req.ConfirmPassword != req.Password {
		ctx.String(http.StatusOK, "两次输入的密码不一致")
	}

	ok, err = u.passwordExp.MatchString(req.Password)
	if err != nil {
		//记录日志
		ctx.String(http.StatusOK, "系统错误")
	}
	if !ok {
		ctx.String(http.StatusOK, "密码必须大于8位，包含数字、特殊字符")
		return
	}

	ctx.String(http.StatusOK, "注册成功")
	fmt.Printf("%v", req)
	//这边就是数据库操作
}

func (u *UserHandler) Login(ctx *gin.Context) {

}

/*func (u *UserHandler) LoginPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", nil)
}*/

func (u *UserHandler) Edit(ctx *gin.Context) {

}

func (u *UserHandler) Profile(ctx *gin.Context) {

}
