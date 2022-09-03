package controller

import "github.com/gin-gonic/gin"

type Controller struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}

func GetControllers() []Controller {
	c := make([]Controller, 10)

	c1 := Controller{
		Method:  "GET",
		Path:    "/",
		Handler: Index,
	}

	//c[0] = c1
	c = append(c, c1)

	return c
}

func Index(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"msg": "hello world.",
	})
}
