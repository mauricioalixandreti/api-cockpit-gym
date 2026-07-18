package service

type User struct {
	Name     string `json:"name" binding:"required,min=6"`
	Email    string `json:"email" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required,min=7"`
}
