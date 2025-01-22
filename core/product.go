package core

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID       uint `gorm:"AUTO_INCREMENT"`
	Name     string
	Quantity int
}
