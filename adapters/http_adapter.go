package adapters

import (
	"strconv"

	"github.com/WarisLi/Golang-mini-project/core"
	"github.com/gofiber/fiber/v2"
)

type HttpProductHandler struct {
	service core.ProductService
}

func NewHttpProductHandler(service core.ProductService) *HttpProductHandler {
	return &HttpProductHandler{service: service}
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
