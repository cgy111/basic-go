package main

import (
	"basic-go/webook/intternal/web"
	"github.com/gin-gonic/gin"
)

func main() {
	//server := gin.Default()
	//
	//u := &web.UserHandler{}
	//
	//server.POST("/users/signup", u.SignUp)
	//
	///*//REST风格
	//server.PUT("/user", func(context *gin.Context) {
	//
	//})*/
	//
	//server.POST("/users/login", u.Login)
	//
	///*//Rest风格
	//server.POST("/users/:id", func(context *gin.Context) {
	//
	//})*/
	//
	////非Rest风格
	//server.POST("/users/edit", u.Edit)
	//
	//server.GET("users/profile", u.Profile)
	//
	///*//REST风格
	//server.GET("users/:id", func(context *gin.Context) {
	//
	//})*/
	//
	//server.Run(":8080")

	//server := web.RegisterRoutes()
	server := gin.Default()
	u := &web.UserHandler{}
	u.RegisterRoutes(server)
	server.Run(":8080")
}
