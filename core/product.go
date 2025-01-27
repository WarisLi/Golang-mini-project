package core

import "gorm.io/gorm"

type Product struct {
	gorm.Model `swaggerignore:"true"`
	ID         uint   `gorm:"AUTO_INCREMENT"`
	Name       string `json:"name" binding:"required" example:"Book"`
	Quantity   int    `json:"quantity" binding:"required" example:"1234"`
}

type ProductInput struct {
	Name     string `json:"name" binding:"required" example:"Book"`
	Quantity int    `json:"quantity" binding:"required" example:"1234"`
}
