package main

import (
	"encoding/json"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func getProducts(db *gorm.DB) ([]ProductModel, error) {
	var products []ProductModel
	result := db.Find(&products)
	print(products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func getProduct(db *gorm.DB, id uint) (*ProductModel, error) {
	var product ProductModel
	result := db.First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func createProduct(db *gorm.DB, product *ProductModel) error {
	result := db.Create(product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func updateProduct(db *gorm.DB, product *ProductModel) error {
	result := db.Model(&product).Updates(product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func deleteProduct(db *gorm.DB, id uint) error {
	var product ProductModel
	result := db.Delete(&product, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func createUser(db *gorm.DB, user *UserModel) error {
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

func loginUser(db *gorm.DB, requestUser *UserModel) error {
	var user UserModel
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
