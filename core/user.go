package core

import "gorm.io/gorm"

type User struct {
	gorm.Model `swaggerignore:"true"`
	ID         uint   `gorm:"AUTO_INCREMENT" json:"-" `
	Username   string `gorm:"unique" json:"username" binding:"required" example:"admin"`
	Password   string `json:"password" binding:"required" example:"Pass@1234"`
}

type UsernamePassword struct {
	Username string `json:"username" binding:"required" example:"admin" validate:"required"`
	Password string `json:"password" binding:"required" example:"Pass@1234" validate:"required"`
}

type LoginSuccess struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type MessageResponse struct {
	Message string `json:"message"`
}
