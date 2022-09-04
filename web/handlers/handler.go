package handlers

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("mustBig", mustBig)
	}
}

/**-----简单请示参数---------*/
func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "hello gin",
	})
}

func PathGet(c *gin.Context) {
	id := c.Param("id")
	user := c.DefaultQuery("user", "zhangsan")
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "hello gin",
		"id":   id,
		"user": user,
	})
}

func PathPost(c *gin.Context) {
	log.Println("-------path post request----------")
	user := c.DefaultPostForm("user", "zhangsan")
	pwd := c.PostForm("pwd")

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "hello gin",
		"pwd":  pwd,
		"user": user,
	})
}

func PathPut(c *gin.Context) {
	user := c.DefaultPostForm("user", "zhangsan")
	age := c.DefaultPostForm("age", "23")
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "hello gin",
		"age":  age,
		"user": user,
	})
}

func PathDelete(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "hello gin",
		"id":   id,
	})
}

/**-----bind请示参数---------*/
type PostParams struct {
	Name string `json:"name" form:"name" binding:"required"`
	Age  int    `json:"age"  form:"age"  binding:"mustBig"`
	Sex  bool   `json:"sex"  form:"sex"`
}

func mustBig(f1 validator.FieldLevel) bool {
	if f1.Field().Interface().(int) <= 18 {
		return false
	}
	return true
}

func PathPostJsonBind(c *gin.Context) {
	var p PostParams
	err := c.ShouldBindJSON(&p)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "参数绑定失败",
			"data": gin.H{},
		})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "成功",
			"data": p,
		})
	}
}

func PathPostFormBind(c *gin.Context) {
	var p PostParams
	err := c.ShouldBindQuery(&p)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "参数绑定失败",
			"data": gin.H{},
		})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "成功",
			"data": p,
		})
	}
}

//`sss`
func PathUpload(c *gin.Context) {
	file, _ := c.FormFile("file")
	fmt.Println("file.size=", file.Size)
	name := c.PostForm("name")
	c.SaveUploadedFile(file, "./"+file.Filename)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "成功",
		"data": "上传文件成功:" + name,
	})
}
