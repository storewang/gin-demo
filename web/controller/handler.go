package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/storewang/gin-demo/web"
)

func init() {
	fmt.Println("----home controller init-----------")
	web.Register(&Home{})
}

type Home struct {
}

func (h *Home) Index_get(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"msg": "hello world.",
	})
}

func (h *Home) SayHello(ctx *gin.Context) {
	s := ctx.DefaultQuery("name", "world.")

	ctx.JSON(200, gin.H{
		"msg": "hello " + s,
	})
}
