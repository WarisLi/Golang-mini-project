package core

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `gorm:"AUTO_INCREMENT"`
	Username string `gorm:"unique"`
	Password string
}

type UserLogin struct {
	Username string `json:"username" binding:"required" example:"admin"`
	Password string `json:"password" binding:"required" example:"Pass@1234"`
}

type LoginSuccess struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type MessageResponse struct {
	Message string `json:"message"`
}
