package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/storewang/gin-demo/infrastructure/util"
	"github.com/storewang/gin-demo/web"
)

func init() {
	fmt.Println("----home controller start register-----------")
	// 自动注册controller中方法为路由(方法必须为GET,POST,PUT,DELETE开头)
	homeHandler := &Home{}
	web.Register("home", homeHandler)

	// 手动注册,用于设置特殊的path
	myHandler := util.GetMethod("SelfHandler", homeHandler)
	route := web.NewRoute("/home/test/:id", "GET", myHandler)
	web.RegisterRoute(route)

	fmt.Println("----home controller register finished-----------")
}

type Home struct {
}

func (h *Home) SelfHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	ctx.JSON(200, gin.H{
		"msg":  "需要要手动注册的",
		"data": id,
	})
}

func (h *Home) Getindex(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"msg": "hello world.",
	})
}

func (h *Home) GetsayHello(ctx *gin.Context) {
	s := ctx.DefaultQuery("name", "world.")

	ctx.JSON(200, gin.H{
		"msg": "hello " + s,
	})
}

func (h *Home) PostHello(ctx *gin.Context) {
	s := ctx.DefaultQuery("name", "world.")

	ctx.JSON(200, gin.H{
		"msg": "hello " + s,
	})
}
