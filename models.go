package main

import "gorm.io/gorm"

type ProductModel struct {
	gorm.Model
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type UserModel struct {
	gorm.Model
	ID       int    `gorm:"AUTO_INCREMENT"`
	Username string `gorm:"unique"`
	Password string
}
