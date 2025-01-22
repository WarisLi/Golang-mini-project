package main

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func createUserHandler(db *gorm.DB, c *fiber.Ctx) error {
	user := new(UserModel)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err := createUser(db, user)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(user)
}

func loginHandler(db *gorm.DB, c *fiber.Ctx) error {
	requestUser := new(UserModel)

	if err := c.BodyParser(requestUser); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err := loginUser(db, requestUser)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"username": requestUser.Username,
		"role":     "admin",
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    t,
		Expires:  time.Now().Add(time.Hour * 72),
		HTTPOnly: true,
	})

	return c.JSON(fiber.Map{
		"message": "Login success",
		"token":   t,
	})
}
