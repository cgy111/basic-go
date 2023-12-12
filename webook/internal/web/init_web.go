package web

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes() *gin.Engine {
	server := gin.Default()

	registerUsersRoutes(server)
	return server
}

func registerUsersRoutes(server *gin.Engine) {
	u := &UserHandler{}
	server.POST("/users/signup", u.Signup)

	/*
		//REST风格
		server.PUT("/user", func(ctx *gin.Context) {

		})
	*/

	server.POST("/users/login", u.Login)

	server.POST("/users/edit", u.Edit)

	/*
		//REST风格
		server.POST("/users/:id", func(ctx *gin.Context) {

		})
	*/

	//server.GET("/users/profilre", u.Profile)

	/*
		//REST风格
		server.GET("/users/:id", func(ctx *gin.Context) {

		})
	*/

}
