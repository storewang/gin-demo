package handlers

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shitou/go-demo-gin/infrastructure/respository"
	"gorm.io/gorm"
)

type UserHandler struct {
	userDao *respository.UserRepository
}

func NewUserHandler(db *gorm.DB) (h *UserHandler) {
	h = &UserHandler{}
	h.init(db)

	return
}
func (u *UserHandler) init(db *gorm.DB) {
	u.userDao = respository.NewUserRepository(db)

	// 初始化数据
	// user1 := po.User{Name: "Jinzhu", Age: 18, Sex: 1}
	// user2 := po.User{Name: "张三", Age: 20, Sex: 1}
	// user3 := po.User{Name: "李四", Age: 30, Sex: 1}

	// result, _ := u.userDao.SaveUser(&user1)
	// fmt.Println("添加用户1：", result)
	// result, _ = u.userDao.SaveUser(&user2)
	// fmt.Println("添加用户2：", result)
	// result, _ = u.userDao.SaveUser(&user3)
	// fmt.Println("添加用户3：", result)
}

func (u *UserHandler) FindById(c *gin.Context) {
	id := c.Param("id")
	userId, _ := strconv.ParseUint(id, 10, 64)
	user := u.userDao.FindUserById(uint(userId))
	fmt.Println(fmt.Sprint("查询用户 id=", userId, ",user=", user))
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": user,
	})
}
