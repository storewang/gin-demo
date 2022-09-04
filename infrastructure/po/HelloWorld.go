package po

type User struct {
	Name string `gorm:"coumn:user_name;type:varchar(100);comment:姓名"`
	Sex  int    `gorm:"coumn:sex;type:tinyint;comment:姓别"`
	Age  int16  `gorm:"coumn:age;type:int;comment:年龄"`
	ID   uint   `gorm:"primarykey;comment:主键"`
}

func (u User) TableName() string {
	return "t_user"
}
