package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/web"
)

func main() {
	ginRouter := gin.Default()
	ginRouter.Handle("GET", "/user", func(ctx *gin.Context) {
		ctx.String(200, "user api")
	})
	ginRouter.Handle("GET", "/news", func(ctx *gin.Context) {
		ctx.String(200, "news api")
	})
	server := web.NewService(
		web.Address(":8002"),
		web.Handler(ginRouter),
	)

	server.Run()
}
