package main

import (
	"github.com/storewang/gin-demo/web"
	_ "github.com/storewang/gin-demo/web/controller"
)

func main() {
	// fmt.Println("hello world.")
	// e := gin.Default()
	// e.GET("/", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"msg": "hello world.",
	// 	})
	// })

	// e.Run(":9999")
	s := web.NewServer()
	s.Run()
}
