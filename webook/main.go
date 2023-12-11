package main

import (
	"basic-go/webook/internal/repository"
	"basic-go/webook/internal/repository/dao"
	"basic-go/webook/internal/service"
	"basic-go/webook/internal/web"
	"basic-go/webook/internal/web/middleware"
	"basic-go/webook/pkg/ginx/middleware/ratelimit"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	_ "github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	redis "github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"strings"
	"time"
)

func main() {
	/*db := initDB()
	server := initWebServer()
	u := initUser(db)
	//u.RegisterRoutesV1(server.Group("/users"))
	u.RegisterRoutes(server)*/

	server := gin.Default()
	server.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "你好，欢迎你")
	})
	server.Run(":8080")

}

func initWebServer() *gin.Engine {
	server := gin.Default()

	//限流
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6319",
	})
	server.Use(ratelimit.NewBuilder(redisClient, time.Second, 100).Build())

	store := memstore.NewStore([]byte("8b8d2e454737a253e0b12365a1ab97e2"), []byte("d9d4e451c8d6b858ca341a3e623411e8"))
	//跨域
	server.Use(cors.New(cors.Config{
		//AllowOrigins: []string{"http://localhost:3000"},
		//AllowMethods: []string{"POST", "GET"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
		//不加，前端拿不到
		ExposeHeaders: []string{"x-jwt-token"},
		//是否允许你带cookie之类的东西
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				//开发环境
				return true
			}
			return strings.Contains(origin, "qiniu.io")
		},
		MaxAge: 12 * time.Hour,
	}))
	//session
	//store := cookie.NewStore([]byte("secret"))
	//store := memstore.NewStore([]byte("8b8d2e454737a253e0b12365a1ab97e2"), []byte("d9d4e451c8d6b858ca341a3e623411e8"))

	/*store, err := redis.NewStore(16, "tcp", "localhost:6379", "",
		[]byte("8b8d2e454737a253e0b12365a1ab97e2"), []byte("d9d4e451c8d6b858ca341a3e623411e8"))

	if err != nil {
		panic(err)
	}*/

	server.Use(sessions.Sessions("mysession", store))

	/*	server.Use(middleware.NewLoginMiddlewareBuilder().
		IgnorePaths("/users/signup").
		IgnorePaths("/users/login").Build())
	*/

	server.Use(middleware.NewLoginJWTMiddlewareBuilder().
		IgnorePaths("/users/signup").
		IgnorePaths("/users/login").Build())

	return server
}

func initUser(db *gorm.DB) *web.UserHandler {
	ud := dao.NewUserDAO(db)
	repo := repository.NewUserRepository(ud)
	svc := service.NewUserService(repo)
	u := web.NewUserHandler(svc)
	return u
}

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13316)/webook?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		//只在初始化的时候panic
		//panic相当于整个goroutine结束
		//一旦初始化过程出错，应用就不要启动了
		panic(err)
	}

	err = dao.InitTable(db)
	if err != nil {
		panic(err)
	}

	return db

}
