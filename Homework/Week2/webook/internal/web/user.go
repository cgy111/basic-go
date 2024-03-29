package web

import (
	"basic-go/Homework/Week2/webook/internal/domain"
	"basic-go/Homework/Week2/webook/internal/service"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 定义和用户有关的路由
type UserHandler struct {
	svc            *service.UserService
	emailExp       *regexp.Regexp
	passwordExp    *regexp.Regexp
	nameExp        *regexp.Regexp
	birthdayExp    *regexp.Regexp
	descriptionExp *regexp.Regexp
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	const (
		emailRegexPattern       = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
		passwordRegexPattern    = "^(?=.*[A-Za-z])(?=.*\\d)(?=.*[$@$!%*#?&])[A-Za-z\\d$@$!%*#?&]{8,}$"
		nameRegexPattern        = "^.{3,16}$"
		birthdayRegexPattern    = "^\\d{4}-\\d{2}-\\d{2}$"
		descriptionRegexPattern = "^.{1,500}$"
	)
	emailExp := regexp.MustCompile(emailRegexPattern, regexp.None)
	passwordExp := regexp.MustCompile(passwordRegexPattern, regexp.None)
	nameEpx := regexp.MustCompile(nameRegexPattern, regexp.None)
	birthdayExp := regexp.MustCompile(birthdayRegexPattern, regexp.None)
	userProfilEpx := regexp.MustCompile(descriptionRegexPattern, regexp.None)
	return &UserHandler{
		svc:            svc,
		emailExp:       emailExp,
		passwordExp:    passwordExp,
		nameExp:        nameEpx,
		birthdayExp:    birthdayExp,
		descriptionExp: userProfilEpx,
	}
}

/*func (u *UserHandler) RegisterRoutesV1(ug *gin.RouterGroup) {
	ug.GET("/profile", u.Profile)
	ug.POST("/signup", u.Signup)
	ug.POST("/login", u.Login)
	ug.POST("/edit", u.Edit)
}*/

func (u *UserHandler) RegisterRoutes(server *gin.Engine) {
	ug := server.Group("/users")
	ug.GET("/profile", u.Profile)
	ug.POST("/signup", u.Signup)
	ug.POST("/login", u.Login)
	ug.POST("/edit", u.Edit)
}

func (u *UserHandler) Signup(ctx *gin.Context) {
	type SignupReq struct {
		Email           string `json:"email"`
		ConfirmPassword string `json:"confirmPassword"`
		Password        string `json:"password"`
	}
	var req SignupReq
	//Bind方法会根据Content-Type 来解析你的数据到req里面
	//解析错了，就会直接写会一个400的错误
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
		ctx.String(http.StatusOK, "你的邮箱格式错误")
	}

	//密码校验
	if req.ConfirmPassword != req.Password {
		ctx.String(http.StatusOK, "两次输入的密码不一致")
	}

	ok, err = u.passwordExp.MatchString(req.Password)
	if err != nil {
		//记录日志
		ctx.String(http.StatusOK, "系统错误")
		return
	}

	if !ok {
		ctx.String(http.StatusOK, "密码必须大于8位，包含数字、特殊字符")
		return
	}

	err = u.svc.SignUp(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})

	if err == service.ErrUserDuplicateEmail {
		ctx.String(http.StatusOK, "邮箱冲突")
		return
	}

	if err != nil {
		ctx.String(http.StatusOK, "系统异常")
		return
	}

	ctx.String(http.StatusOK, "注册成功")
	//fmt.Printf("%v", req)

}

func (u *UserHandler) Login(ctx *gin.Context) {
	type LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req LoginReq

	if err := ctx.Bind(&req); err != nil {
		return
	}

	user, err := u.svc.Login(ctx, req.Email, req.Password)
	if err == service.ErrInvalidUserOrPassword {
		ctx.String(http.StatusOK, "用户名或密码错误")
		return
	}
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}

	//在这里登录成功了
	//设置session
	sess := sessions.Default(ctx)
	//我可以随便设置值了
	//你要放在session里面的值
	sess.Set("userId", user.Id)
	sess.Save()
	ctx.String(http.StatusOK, "登录成功")
	return
}

func (u *UserHandler) Edit(ctx *gin.Context) {
	type EditReq struct {
		Id          int64  `json:"id"`
		Name        string `json:"name"`
		Birthday    string `json:"birthday"`
		Description string `json:"description"`
	}
	var req EditReq

	if err := ctx.Bind(&req); err != nil {
		return
	}

	//验证昵称
	ok, err := u.nameExp.MatchString(req.Name)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	if !ok {
		ctx.String(http.StatusOK, "昵称的长度在3到16个字符之间")
		return
	}

	//验证生日格式
	ok, err = u.birthdayExp.MatchString(req.Birthday)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}

	if !ok {
		ctx.String(http.StatusOK, "请输入如\"1992-01-01\"这种格式的生日日期")
		return
	}

	ok, err = u.descriptionExp.MatchString(req.Description)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}

	if !ok {
		ctx.String(http.StatusOK, "个人简介的长度在1到500个字符之间")
		return
	}
	err = u.svc.Edit(ctx, domain.User{
		Id:          req.Id,
		Name:        req.Name,
		Birthday:    req.Birthday,
		Description: req.Description,
	})
	if err != nil {
		ctx.String(http.StatusOK, "系统异常")
		return
	}
	//ctx.String(http.StatusOK, "修改成功")
	ctx.JSON(http.StatusOK, gin.H{"message": "用户信息修改成功"})
}

func (u *UserHandler) Profile(ctx *gin.Context) {
	//ctx.String(http.StatusOK, "这是你的profile")
	userIdStr := ctx.Query("id")

	if userIdStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id不能为空"})
		return
	}

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id不正确"})
		return
	}
	mes, err := u.svc.Profile(ctx, userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "查询失败"})
		return
	}

	type UserResponse struct {
		Id          int64  `json:"id"`
		Email       string `json:"email"`
		Name        string `json:"name"`
		Birthday    string `json:"birthday"`
		Description string `json:"description"`
	}

	mess := UserResponse{
		Id:          mes.Id,
		Email:       mes.Email,
		Name:        mes.Name,
		Birthday:    mes.Birthday,
		Description: mes.Description,
	}
	ctx.JSON(http.StatusOK, mess)
}
