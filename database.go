package main

import (
	"encoding/json"
	"fmt"

	"github.com/WarisLi/Golang-mini-project/core"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func createUser(db *gorm.DB, user *core.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	result := db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func loginUser(db *gorm.DB, requestUser *core.User) error {
	var user core.User
	result := db.Where("username = ?", requestUser.Username).First(&user)
	if result.Error != nil {
		return result.Error
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestUser.Password))
	if err != nil {
		return err
	}
	json, _ := json.Marshal(requestUser)
	fmt.Println(string(json))
	return nil
}
