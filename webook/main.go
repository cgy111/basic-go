package main

import (
	"basic-go/webook/intternal/web"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func main() {

	server := gin.Default()

	server.Use(func(ctx *gin.Context) {
		println("这是第一个middleware")

	})

	server.Use(func(ctx *gin.Context) {
		println("这是第二个middleware")
	})

	server.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		//AllowMethods: []string{"PUT", "PATCH"},
		AllowHeaders: []string{"content-type", "authorization"},
		//ExposeHeaders:    []string{"Content-Length"},
		//是否允许带cooike
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			//return origin == "https://localhost:8081"
			//if strings.Contains(origin, "localhost") {
			if strings.HasPrefix(origin, "http://localhost") {
				//	开发环境
				return true
			}
			return strings.Contains(origin, ":8081")
		},
		MaxAge: 12 * time.Hour,
	}))

	//u := &web.UserHandler{}
	u := web.NewUserHandler()
	//u.RegisterRoutesV1(server.Group("/users"))
	u.RegisterRoutes(server)
	server.Run(":8081")
}
