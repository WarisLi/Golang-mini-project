package core

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `gorm:"AUTO_INCREMENT"`
	Username string `gorm:"unique"`
	Password string
}
