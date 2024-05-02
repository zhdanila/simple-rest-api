package models

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignInInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
