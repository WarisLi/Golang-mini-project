package config

import (
	"fmt"

	"github.com/WarisLi/Golang-mini-project/internal/core/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host         = "localhost"
	port         = 5432
	databaseName = "mydatabase"
	username     = "myuser"
	password     = "mypassword"
)

func initData(db *gorm.DB) {
	// init test data
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("Pass@12345"), bcrypt.DefaultCost)
	user := models.User{
		Username: "user_1",
		Password: string(hashedPassword),
	}
	if result := db.Create(&user); result.Error != nil {
		fmt.Printf("Initial user data failed %s\n", result.Error)
	}

	books := []*models.Product{{
		Name:     "Book A",
		Quantity: 1200,
	}, {
		Name:     "Book B",
		Quantity: 400,
	},
	}
	if result := db.Create(&books); result.Error != nil {
		fmt.Printf("Initial product data failed %s\n", result.Error)
	}

	fmt.Printf("Initial data completed\n")
}

func SetupDB() *gorm.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, databaseName)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{TranslateError: true,
		Logger: logger.Default.LogMode(logger.Silent)})

	if err != nil {
		panic("fail to connect database\n")
	}

	fmt.Printf("Database Connecction successful\n")

	models := []interface{}{models.Product{}, models.User{}}

	db.AutoMigrate(models...)
	fmt.Printf("Database migration completed\n")

	for _, model := range models {
		if err := db.Session(&gorm.Session{AllowGlobalUpdate: true}).Unscoped().Delete(model).Error; err != nil {
			fmt.Printf("Data reset failed %s\n", err.Error())
		}
	}
	fmt.Printf("Data reset completed\n")

	initData(db)

	return db
}
