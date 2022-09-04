package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shitou/go-demo-gin/applications"
	"github.com/shitou/go-demo-gin/applications/dto"
	"github.com/shitou/go-demo-gin/infrastructure/respository"
	"github.com/shitou/go-demo-gin/modles"
	"github.com/shitou/go-demo-gin/web/vo"
	"gorm.io/gorm"
)

type UserHandler struct {
	userAppSvc *applications.UserAppService
}

func NewUserHandler(db *gorm.DB) (h *UserHandler) {
	h = &UserHandler{}
	h.init(db)

	return
}
func (u *UserHandler) init(db *gorm.DB) {
	userService := modles.NewUserHandler(respository.NewUserRepository(db))
	userAppService := applications.NewUserAppService(userService)
	u.userAppSvc = userAppService
}

func (u *UserHandler) FindById(c *gin.Context) {
	id := c.Param("id")
	userId, _ := strconv.ParseUint(id, 10, 64)
	user := u.userAppSvc.FindUserById(userId)

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": user,
	})
}

func (u *UserHandler) AddUser(c *gin.Context) {
	var user vo.UserVO
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "添加用户失败",
			"data": gin.H{},
		})
	} else {
		userDTO := dto.UserDTO{Name: user.Name, Age: user.Age, Sex: user.Sex}
		uid := u.userAppSvc.AddUser(&userDTO)

		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "添加用户成功",
			"data": uid,
		})
	}

}
