package vo

type UserVO struct {
	Name string `json:"name" form:"name" binding:"required"`
	Age  int16  `json:"age"  form:"age"`
	Sex  int    `json:"sex"  form:"sex"`
	ID   uint   `json:"id"  form:"id"`
}
