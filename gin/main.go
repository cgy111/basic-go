package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	service := gin.Default()
	service.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello,go")
	})
	service.POST("/post", func(context *gin.Context) {
		context.String(http.StatusOK, "hello,post")
	})
	service.GET("/users/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.String(http.StatusOK, "hello,这是参数路由"+name)
	})
	service.GET("/views/*.html", func(context *gin.Context) {
		page := context.Param(".html")
		context.String(http.StatusOK, "hello,这是通配符路由"+page)
	})
	service.GET("/order", func(context *gin.Context) {
		oid := context.Query("id")
		context.String(http.StatusOK, "hello,这是查询参数"+oid)
	})
	service.GET("/items/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello,这是items")
	})
	service.GET("/cgy/*abc", func(context *gin.Context) {
		context.String(http.StatusOK, "hello,这是cgy")
	})
	service.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
