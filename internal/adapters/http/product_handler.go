package http

import (
	"strconv"

	"github.com/WarisLi/Golang-mini-project/internal/core/models"
	"github.com/WarisLi/Golang-mini-project/internal/core/ports"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type HttpProductHandler struct {
	service ports.ProductService
}

func NewHttpProductHandler(service ports.ProductService) *HttpProductHandler {
	return &HttpProductHandler{service: service}
}

// Handler functions
// GetProducts godoc
// @Summary Get all products
// @Description Get details of all products
// @Tags product
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {array} ports.Product
// @Router /product [get]
func (h *HttpProductHandler) GetProducts(c *fiber.Ctx) error {
	// call primary port function
	products, err := h.service.GetProducts()
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.Status(fiber.StatusOK).JSON(products)
}

// Handler functions
// GetProduct godoc
// @Summary Get product
// @Description Get details of product
// @Tags product
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} models.Product
// @Param id path uint true "ID"
// @Router /product/{id} [get]
func (h *HttpProductHandler) GetProduct(c *fiber.Ctx) error {
	productId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.MessageResponse{Message: err.Error()})
	}

	product, err := h.service.GetProduct(uint(productId))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.MessageResponse{Message: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(product)
}

// Handler functions
// CreateProduct godoc
// @Summary Create product
// @Description Create product
// @Tags product
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} models.MessageResponse
// @Param product body models.ProductInput true "Product"
// @Router /product [POST]
func (h *HttpProductHandler) CreateProduct(c *fiber.Ctx) error {
	var product models.ProductInput
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.MessageResponse{Message: err.Error()})
	}

	var validate = validator.New()
	if err := validate.Struct(product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.MessageResponse{Message: err.Error()})
	}

	if err := h.service.CreateProduct(product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.MessageResponse{Message: err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(models.MessageResponse{Message: "success"})
}

// Handler functions
// UpdateProduct godoc
// @Summary Update product
// @Description Update product
// @Tags product
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} models.MessageResponse
// @Param product body models.ProductInput true "Product"
// @Param id path uint true "ID"
// @Router /product/{id} [PUT]
func (h *HttpProductHandler) UpdateProduct(c *fiber.Ctx) error {
	productId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.MessageResponse{Message: err.Error()})
	}

	productUpdate := new(models.ProductInput)
	if err := c.BodyParser(productUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.MessageResponse{Message: err.Error()})
	}

	var validate = validator.New()
	if err := validate.Struct(productUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.MessageResponse{Message: err.Error()})
	}

	if err := h.service.UpdateProduct(uint(productId), *productUpdate); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.MessageResponse{Message: err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(models.MessageResponse{Message: "success"})
}

// Handler functions
// DeleteProduct godoc
// @Summary Delete product
// @Description Delete product
// @Tags product
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} models.MessageResponse
// @Param id path uint true "ID"
// @Router /product/{id} [DELETE]
func (h *HttpProductHandler) DeleteProduct(c *fiber.Ctx) error {
	productId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.MessageResponse{Message: err.Error()})
	}

	err = h.service.DeleteProduct(uint(productId))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.Status(fiber.StatusOK).JSON(models.MessageResponse{Message: "success"})
}
