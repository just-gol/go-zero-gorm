package gorm

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
}

func (User) TableName() string {
	return "user"
}
