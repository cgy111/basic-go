package web

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes() *gin.Engine {
	server := gin.Default()

	registerUsersRoutes(server)

	return server

	//server.Run(":8080")
}

func registerUsersRoutes(server *gin.Engine) {
	u := &UserHandler{}

	server.POST("/users/signup", u.SignUp)

	/*//REST风格
	server.PUT("/user", func(context *gin.Context) {

	})*/

	server.POST("/users/login", u.Login)

	/*//Rest风格
	server.POST("/users/:id", func(context *gin.Context) {

	})*/

	//非Rest风格
	server.POST("/users/edit", u.Edit)

	server.GET("users/profile", u.Profile)

	/*//REST风格
	server.GET("users/:id", func(context *gin.Context) {

	})*/
}
