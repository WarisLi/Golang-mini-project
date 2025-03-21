package http

import (
	"errors"
	"time"

	"github.com/WarisLi/Golang-mini-project/internal/core/models"
	"github.com/WarisLi/Golang-mini-project/internal/core/ports"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type HttpUserHandler struct {
	service ports.UserService
}

func NewHttpUserHandler(service ports.UserService) *HttpUserHandler {
	return &HttpUserHandler{service: service}
}

// Handler functions
// CreateUser godoc
// @Summary Create user
// @Description Create user
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body models.User true "Username/password"
// @Success 201 {object} models.MessageResponse
// @Router /user [post]
func (h *HttpUserHandler) CreateUser(c *fiber.Ctx) error {
	var user models.UsernamePassword
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.MessageResponse{Message: err.Error()})
	}

	var validate = validator.New()
	if err := validate.Struct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.MessageResponse{Message: err.Error()})
	}

	if err := h.service.RegisterUser(user); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return c.Status(fiber.StatusBadRequest).JSON(models.MessageResponse{Message: err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(models.MessageResponse{Message: err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(models.MessageResponse{Message: "success"})
}

// Handler functions
// LoginUser godoc
// @Summary Login user
// @Description Login user
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body models.User true "Username/password"
// @Success 200 {object} models.LoginSuccess
// @Router /user/login [post]
func (h *HttpUserHandler) LoginUser(c *fiber.Ctx) error {
	var requestUser models.UsernamePassword

	if err := c.BodyParser(&requestUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.MessageResponse{Message: err.Error()})
	}

	var validate = validator.New()
	if err := validate.Struct(requestUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.MessageResponse{Message: err.Error()})
	}

	token, err := h.service.LoginUser(requestUser)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(models.MessageResponse{Message: "The username or password is incorrect"})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 72),
		HTTPOnly: true,
	})

	return c.JSON(models.LoginSuccess{Message: "Login success", Token: token})
}
