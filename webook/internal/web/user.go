package web

import (
	"basic-go/webook/internal/domain"
	"basic-go/webook/internal/service"
	ijwt "basic-go/webook/internal/web/jwt"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel/trace"
	"net/http"
)

const biz = "login"

// 确保UserHandler实现了handler接口
// 实现了一个对象
var _ handler = &UserHandler{}

// 这种更优雅
var _ handler = (*UserHandler)(nil)

// 定义和用户有关的路由
type UserHandler struct {
	svc            service.UserService
	codeSvc        service.CodeService
	emailExp       *regexp.Regexp
	passwordExp    *regexp.Regexp
	nameExp        *regexp.Regexp
	birthdayExp    *regexp.Regexp
	descriptionExp *regexp.Regexp
	phoneExp       *regexp.Regexp
	//组合
	ijwt.Handler
	cmd redis.Cmdable
}

func NewUserHandler(svc service.UserService, codeSvc service.CodeService, jwtHdl ijwt.Handler) *UserHandler {
	const (
		emailRegexPattern       = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
		passwordRegexPattern    = "^(?=.*[A-Za-z])(?=.*\\d)(?=.*[$@$!%*#?&])[A-Za-z\\d$@$!%*#?&]{8,}$"
		nameRegexPattern        = "^.{3,16}$"
		birthdayRegexPattern    = "^\\d{4}-\\d{2}-\\d{2}$"
		descriptionRegexPattern = "^.{1,500}$"
		phoneRegexPattern       = "^1[3-9]\\d{9}$"
	)
	emailExp := regexp.MustCompile(emailRegexPattern, regexp.None)
	passwordExp := regexp.MustCompile(passwordRegexPattern, regexp.None)
	nameEpx := regexp.MustCompile(nameRegexPattern, regexp.None)
	birthdayExp := regexp.MustCompile(birthdayRegexPattern, regexp.None)
	userProfilEpx := regexp.MustCompile(descriptionRegexPattern, regexp.None)
	phoneExp := regexp.MustCompile(phoneRegexPattern, regexp.None)
	return &UserHandler{
		svc:            svc,
		codeSvc:        codeSvc,
		emailExp:       emailExp,
		passwordExp:    passwordExp,
		nameExp:        nameEpx,
		birthdayExp:    birthdayExp,
		descriptionExp: userProfilEpx,
		phoneExp:       phoneExp,
		Handler:        jwtHdl,
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
	//ug.GET("/profile", u.Profile)
	ug.GET("/profile", u.ProfileJWT)
	ug.POST("/signup", u.SignUp)
	//ug.POST("/login", u.Login)
	ug.POST("/login", u.LoginJWT)
	ug.POST("/logout", u.LogoutJWT)
	ug.POST("/edit", u.Edit)
	ug.POST("/login_sms/code/send", u.SendLoginSmsCode)
	ug.POST("/login_sms", u.LoginSMS)
	ug.POST("/refresh_token", u.RefreshToken)
}

func (u *UserHandler) LogoutJWT(ctx *gin.Context) {
	err := u.ClearToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, Result{
			Code: 5,
			Msg:  "退出登录失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, Result{
		Msg: "退出登录成功",
	})

}

// RefreshToken 可以同时刷新长短 token,用redis 来记录是否有效,既refresh_token 是一次性的
// 参考登录校验部分,比较 User-Agent 来增强安全性
func (u *UserHandler) RefreshToken(ctx *gin.Context) {
	//只有这个接口,拿出来的才是refresh token,其他地方都是access token
	refreshToken := u.ExtractToken(ctx)
	var rc ijwt.RefreshClaims
	token, err := jwt.ParseWithClaims(refreshToken, &rc, func(token *jwt.Token) (interface{}, error) {
		return ijwt.RtKey, nil
	})
	if err != nil || !token.Valid {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	err = u.CheckSession(ctx, rc.Ssid)
	if err != nil {
		//要么redis有问题，要么已经退出登录
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	//搞个新的access token
	u.SetJWTToken(ctx, rc.Uid, rc.Ssid)
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	ctx.JSON(http.StatusOK, Result{
		Msg: "刷新成功",
	})
}

func (u *UserHandler) LoginSMS(ctx *gin.Context) {
	type Req struct {
		Phone string `json:"phone"`
		Code  string `json:"code"`
	}
	var req Req
	if err := ctx.Bind(&req); err != nil {
		return
	}
	ok, err := u.codeSvc.Verify(ctx, biz, req.Phone, req.Code)
	if err != nil {
		ctx.JSON(http.StatusOK, Result{
			Code: 5,
			Msg:  "系统错误",
		})
		return
	}
	if !ok {
		ctx.JSON(http.StatusOK, Result{
			Code: 4,
			Msg:  "验证码有误",
		})
		return
	}
	//手机号码会不会是一个新用户
	//
	user, err := u.svc.FindOrCreate(ctx, req.Phone)
	if err != nil {
		ctx.JSON(http.StatusOK, Result{
			Code: 5,
			Msg:  "系统错误",
		})
		return
	}

	//id从哪里来？
	if err = u.SetLoginToken(ctx, user.Id); err != nil {
		//记录日志
		ctx.JSON(http.StatusOK, Result{
			Code: 5,
			Msg:  "系统错误",
		})
		return
	}
	//fmt.Println("ll")
	ctx.JSON(http.StatusOK, Result{
		//Code: 2,
		Msg: "验证码验证通过",
	})
	//fmt.Println("dd")
}

func (u *UserHandler) SendLoginSmsCode(ctx *gin.Context) {
	//fmt.Println("send sms code")
	type Req struct {
		Phone string `json:"phone"`
	}

	var req Req
	//fmt.Println(req.Phone)
	if err := ctx.Bind(&req); err != nil {
		return
	}

	//if req.Phone == "" {
	//	ctx.JSON(http.StatusOK, Result{
	//		Code: 4,
	//		Msg:  "输入有误",
	//	})
	//	return
	//}
	ok, err := u.phoneExp.MatchString(req.Phone)
	if err != nil {
		ctx.JSON(http.StatusOK, Result{
			Code: 5,
			Msg:  "系统错误",
		})
		return
	}

	if !ok {
		ctx.JSON(http.StatusOK, Result{
			Code: 4,
			Msg:  "输入有误",
		})
		return
	}

	err = u.codeSvc.Send(ctx, biz, req.Phone)
	switch err {
	case nil:
		ctx.JSON(http.StatusOK, Result{
			Msg: "发送成功",
		})
	case service.ErrorCodeSendTooMany:
		ctx.JSON(http.StatusOK, Result{
			Msg: "发送太频繁，请稍后再试",
		})

	default:
		ctx.JSON(http.StatusOK, Result{
			Code: 5,
			Msg:  "系统错误",
		})
	}
}

func (u *UserHandler) SignUp(ctx *gin.Context) {
	type SignUpReq struct {
		Email           string `json:"email"`
		ConfirmPassword string `json:"confirmPassword"`
		Password        string `json:"password"`
	}
	var req SignUpReq
	//fmt.Println("2")

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
		ctx.String(http.StatusOK, "你的邮箱格式不对")
		return
	}

	//密码校验
	if req.ConfirmPassword != req.Password {
		ctx.String(http.StatusOK, "两次输入的密码不一致")
		return
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

	err = u.svc.SignUp(ctx.Request.Context(), domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	if err == service.ErrUserDuplicate {
		span := trace.SpanFromContext(ctx.Request.Context())
		span.AddEvent("邮箱冲突")
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

func (u *UserHandler) LoginJWT(ctx *gin.Context) {
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

	//步骤2
	//在这里用JWT设置登录态
	//生成一个JWT token
	err = u.SetLoginToken(ctx, user.Id)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}

	ctx.String(http.StatusOK, "登录成功")
	return
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
	sess.Options(sessions.Options{
		//https协议
		//Secure: true,
		//HttpOnly: true,
		MaxAge: 30 * 60,
	})
	sess.Save()
	ctx.String(http.StatusOK, "登录成功")
	return
}

func (u *UserHandler) Logout(ctx *gin.Context) {
	sess := sessions.Default(ctx)
	//我可以随便设置值了
	//你要放在session里面的值
	sess.Options(sessions.Options{
		//https协议
		//Secure: true,
		//HttpOnly: true,
		MaxAge: -1,
	})
	sess.Save()
	ctx.String(http.StatusOK, "退出登录成功")
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

func (u *UserHandler) ProfileJWT(ctx *gin.Context) {
	c, ok := ctx.Get("claims")
	//可以断点，必然有claims
	if !ok {
		//监控住这里
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	//ok代表是不是*UserClaims
	claims, ok := c.(*ijwt.UserClaims)
	if !ok {
		//也可以监控住这里
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	println(claims.Uid)
	ctx.String(http.StatusOK, "你的 profile")
	//这里补充profile的其他代码
}

//func (u *UserHandler) Profile(ctx *gin.Context) {
//	//ctx.String(http.StatusOK, "这是你的profile")
//	userIdStr := ctx.Query("id")
//
//	if userIdStr == "" {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id不能为空"})
//		return
//	}
//
//	userId, err := strconv.Atoi(userIdStr)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id不正确"})
//		return
//	}
//	mes, err := u.svc.Profile(ctx, userId)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": "查询失败"})
//		return
//	}
//
//	type UserResponse struct {
//		Id          int64  `json:"id"`
//		Email       string `json:"email"`
//		Name        string `json:"name"`
//		Birthday    string `json:"birthday"`
//		Description string `json:"description"`
//	}
//
//	mess := UserResponse{
//		Id:          mes.Id,
//		Email:       mes.Email,
//		Name:        mes.Name,
//		Birthday:    mes.Birthday,
//		Description: mes.Description,
//	}
//	ctx.JSON(http.StatusOK, mess)
//}

func (c *UserHandler) Profile(ctx *gin.Context) {
	//ctx.String(http.StatusOK, "这是你的profile")
	//userIdStr := ctx.Query("id")
	//
	//if userIdStr == "" {
	//	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id不能为空"}) {

}
