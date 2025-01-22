package main

import (
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Handler functions
// getProductsHandler godoc
// @Summary Get all products
// @Description Get details of all products
// @Tags products
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {array} Product
// @Router /product [get]
func getProductsHandler(db *gorm.DB, c *fiber.Ctx) error {
	products, err := getProducts(db)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(products)
}

func getProductHandler(db *gorm.DB, c *fiber.Ctx) error {
	productId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	product, err := getProduct(db, uint(productId))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(product)
}

func createProductHandler(db *gorm.DB, c *fiber.Ctx) error {
	product := new(ProductModel)
	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err := createProduct(db, product)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(product)
}

func updateProductHandler(db *gorm.DB, c *fiber.Ctx) error {
	productId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	productUpdate := new(ProductModel)
	if err := c.BodyParser(productUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	productUpdate.ID = uint(productId)

	err = updateProduct(db, productUpdate)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(fiber.Map{"message": "Update successful"})
}

func deleteProductHandler(db *gorm.DB, c *fiber.Ctx) error {
	productId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err = deleteProduct(db, uint(productId))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(fiber.Map{"message": "Delete successful"})
}

func uploadFileHandler(c *fiber.Ctx) error {
	file, err := c.FormFile("image")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err = c.SaveFile(file, os.Getenv("UPLOAD_FOLDER_PATH")+file.Filename)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendString("File upload complete!")
}
