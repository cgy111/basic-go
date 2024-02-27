package main

import (
	"bytes"
	_ "github.com/gin-contrib/sessions/redis"
	"github.com/spf13/viper"
)

func main() {
	//db := initDB()
	//rdb := initRedis()
	//server := initWebServer()
	//u := initUser(db, rdb)
	////u.RegisterRoutesV1(server.Group("/users"))
	//u.RegisterRoutes(server)
	//
	//server := gin.Default()
	//server.GET("/hello", func(ctx *gin.Context) {
	//	ctx.String(http.StatusOK, "你好，欢迎你")
	//})
	//initViperV1()
	initViperReader()
	server := InitWebServer()

	server.Run(":8080")

}

// 开发环境及联调环境
func initViperReader() {
	viper.SetConfigType("yaml")

	cfg := `
db.mysql:
  dsn:  "root:root@tcp(47.123.5.217:13316)/webook"

redis:
  addr: "47.123.5.217:6379"
`

	err := viper.ReadConfig(bytes.NewReader([]byte(cfg)))
	if err != nil {
		panic(err)
	}
}

func initViperV1() {
	//viper.SetDefault("db.mysql.dsn",
	//	"root:root@tcp(47.123.5.217:13316)/mysql")
	viper.SetConfigFile("config/dev.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initViper() {
	//配置文件的名字，但是不包含文件扩展名
	//不包含 .go , .yaml 之类的后缀
	viper.SetConfigName("dev")
	//告诉viper配置文件是yaml格式
	//现实中，用很多格式，json,yaml,toml,xml，init
	viper.SetConfigType("yaml")
	//当前工作目录下的 config 子目录
	viper.AddConfigPath("./config")
	//viper.AddConfigPath("/tmp/config")
	//viper.AddConfigPath("/ect/webook")
	//读取配置到 viper 里面,或者可以理解为加载到内存里面
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	//otherViper := viper.New()
	//otherViper.SetConfigName("myjson")
	//otherViper.AddConfigPath("./config")
	//otherViper.SetConfigType("json")
}

/*func initWebServer() *gin.Engine {
server := gin.Default()

//限流
//redisClient := redis.NewClient(&redis.Options{
//	Addr: config.Config.Redis.Addr,
//})
//server.Use(ratelimit.NewBuilder(redisClient, time.Second, 100).Build())

//store := memstore.NewStore([]byte("8b8d2e454737a253e0b12365a1ab97e2"), []byte("d9d4e451c8d6b858ca341a3e623411e8"))
//跨域
server.Use(cors.New(cors.Config{
	AllowOrigins: []string{"http://localhost:3000"},
	AllowMethods: []string{"POST", "GET"},
	AllowHeaders: []string{"Content-Type", "Authorization"},
	//不加，前端拿不到
	ExposeHeaders: []string{"x-jwt-token"},
	//是否允许你带cookie之类的东西
	AllowCredentials: true,
	AllowOriginFunc: func(origin string) bool {
		if strings.HasPrefix(origin, "http://localhost:") {
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

//server.Use(sessions.Sessions("mysession", store))

/*	server.Use(middleware.NewLoginMiddlewareBuilder().
		IgnorePaths("/users/signup").
		IgnorePaths("/users/login").Build())


	server.Use(middleware.NewLoginJWTMiddlewareBuilder().
		IgnorePaths("/users/signup").
		IgnorePaths("/users/login_sms/code/send").
		IgnorePaths("/users/login_sms").
		IgnorePaths("/users/login").Build())

	return server
}
*/
/*func initUser(db *gorm.DB, rdb redis.Cmdable) *web.UserHandler {
	ud := dao.NewUserDAO(db)
	//c := cache.NewUserCache(redis.NewClient(&redis.Options{
	//	Addr: config.Config.Redis.Addr,
	//}))
	uc := cache.NewUserCache(rdb)
	repo := repository.NewUserRepository(ud, uc)
	svc := service.NewUserService(repo)
	codeCache := cache.NewCodeCache(rdb)
	codeRepo := repository.NewCodeRepository(codeCache)
	smsSvc := memory.NewService()
	codeSvc := service.NewCodeService(codeRepo, smsSvc)
	u := web.NewUserHandler(svc, codeSvc)
	return u
}*/

/*
被wire改造
func initDB() *gorm.DB {
	//fmt.Println("初始化数据库")
	db, err := gorm.Open(mysql.Open(config.Config.DB.DSN))
	fmt.Println(err)
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
}*/

/*
//被wire改造
func initRedis() redis.Cmdable {
	redisClient := redis.NewClient(&redis.Options{
		Addr: config.Config.Redis.Addr,
	})
	return redisClient
}*/
