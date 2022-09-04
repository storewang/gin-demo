package respository

import (
	"fmt"

	"github.com/shitou/go-demo-gin/infrastructure/po"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindUserById(id uint) (u *po.User) {
	user := po.User{}
	r.db.Where("id = ?", id).First(&user)

	fmt.Println(fmt.Sprint("用户 id=", id, ",user=", user))
	return &user
}

func (r *UserRepository) SaveUser(u *po.User) (uint, error) {
	d := r.db.Create(u)
	if d.Error != nil {
		return 0, d.Error
	}

	return u.ID, nil
}
