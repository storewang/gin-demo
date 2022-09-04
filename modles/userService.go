package modles

import (
	"fmt"
	"strconv"

	"github.com/shitou/go-demo-gin/applications/dto"
	"github.com/shitou/go-demo-gin/infrastructure/po"
	"github.com/shitou/go-demo-gin/infrastructure/respository"
)

type UserService struct {
	userDao *respository.UserRepository
}

func NewUserHandler(userDao *respository.UserRepository) (svc *UserService) {
	svc = &UserService{userDao: userDao}
	return
}

func (userSvc *UserService) AddUser(user *dto.UserDTO) int {
	userPO := po.User{Name: user.Name, Age: user.Age, Sex: user.Sex}
	userId, err := userSvc.userDao.SaveUser(&userPO)
	if err != nil {
		fmt.Println("添加用户信息失败：", user)
	}

	idstr := strconv.FormatUint(uint64(userId), 10)
	id, _ := strconv.Atoi(idstr)
	return id
}

func (userSvc *UserService) FindUserById(id uint64) *dto.UserDTO {
	user := userSvc.userDao.FindUserById(uint(id))

	return &dto.UserDTO{Name: user.Name, Age: user.Age, ID: user.ID, Sex: user.Sex}
}
