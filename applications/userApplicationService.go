package applications

import (
	"github.com/shitou/go-demo-gin/applications/dto"
	"github.com/shitou/go-demo-gin/modles"
)

type UserAppService struct {
	svc *modles.UserService
}

func NewUserAppService(svc *modles.UserService) *UserAppService {
	userAppService := &UserAppService{svc: svc}
	return userAppService
}

func (appsvc *UserAppService) AddUser(user *dto.UserDTO) int {
	return appsvc.svc.AddUser(user)
}

func (appsvc *UserAppService) FindUserById(id uint64) *dto.UserDTO {

	return appsvc.svc.FindUserById(id)
}
