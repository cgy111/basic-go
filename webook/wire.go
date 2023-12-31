//go:build wireinject

package main

import "github.com/gin-gonic/gin"

func InitWebServer() *gin.Engine {
	wire.Build()
}
