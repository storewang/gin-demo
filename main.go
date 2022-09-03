package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello world.")
	e := gin.Default()
	e.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"msg": "hello world.",
		})
	})

	e.Run(":9999")
}
