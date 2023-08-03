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
	service.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
