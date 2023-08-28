package models

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type User struct {
	Value string `json:"value" binding:"required"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}
