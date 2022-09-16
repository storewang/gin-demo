package repository

import "github.com/shitou/go-demo-gin/infrastructure/po"

type IUserRepository interface {
	FindUserById(id uint) (u *po.User)
	SaveUser(u *po.User) (uint, error)
}
