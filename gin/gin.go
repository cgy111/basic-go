package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()
	server.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello GO!")
	})

	go func() {
		server1 := gin.Default()
		server1.GET("/hello", func(c *gin.Context) {
			c.String(http.StatusOK, "Hello GO!")
		})
		server1.Run(":8081") // 监听并在 0.0.0.0:8081 上启动服务
	}()

	server.POST("/post", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello post")
	})

	server.GET("/users/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.String(http.StatusOK, "hello , 这是参数路由"+name)
	})

	server.GET("/views/*.html", func(ctx *gin.Context) {
		page := ctx.Param(".html")
		ctx.String(http.StatusOK, "hello,这是通配符路由"+page)
	})

	server.GET("/order", func(ctx *gin.Context) {
		oid := ctx.Query("id")
		ctx.String(http.StatusOK, "hello,这是查询参数"+oid)
	})

	//server.GET("/items/", func(ctx *gin.Context) {
	//	ctx.String(http.StatusOK, "hello,这是itmes")
	//})

	server.GET("/items/*abc", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello,这是itmes")
	})

	server.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务

}
