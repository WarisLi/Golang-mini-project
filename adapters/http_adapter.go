package adapters

import (
	"os"
	"strconv"
	"time"

	"github.com/WarisLi/Golang-mini-project/core"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// connect primary port
type HttpProductHandler struct {
	service core.ProductService
}

type HttpUserHandler struct {
	service core.UserService
}

func NewHttpProductHandler(service core.ProductService) *HttpProductHandler {
	return &HttpProductHandler{service: service}
}

func NewHttpUserHandler(service core.UserService) *HttpUserHandler {
	return &HttpUserHandler{service: service}
}

// Handler functions
// getProductsHandler godoc
// @Summary Get all products
// @Description Get details of all products
// @Tags product
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {array} core.Product
// @Router /product [get]
func (h *HttpProductHandler) GetProducts(c *fiber.Ctx) error {
	// call primary port function
	products, err := h.service.GetProducts()
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(products)
}

func (h *HttpProductHandler) GetProduct(c *fiber.Ctx) error {
	productId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	product, err := h.service.GetProduct(uint(productId))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.JSON(product)
}

func (h *HttpProductHandler) CreateProduct(c *fiber.Ctx) error {
	var product core.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if err := h.service.CreateProduct(product); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(product)
}

func (h *HttpProductHandler) UpdateProduct(c *fiber.Ctx) error {
	productId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	productUpdate := new(core.Product)
	if err := c.BodyParser(productUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	productUpdate.ID = uint(productId)

	if err := h.service.UpdateProduct(*productUpdate); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(productUpdate)
}

func (h *HttpProductHandler) DeleteProduct(c *fiber.Ctx) error {
	productId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err = h.service.DeleteProduct(uint(productId))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(fiber.Map{"message": "Delete successful"})
}

func (h *HttpUserHandler) CreateUser(c *fiber.Ctx) error {
	var user core.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if err := h.service.RegisterUser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func (h *HttpUserHandler) LoginUser(c *fiber.Ctx) error {
	requestUser := new(core.User)

	if err := c.BodyParser(requestUser); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err := h.service.LoginUser(*requestUser)
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
