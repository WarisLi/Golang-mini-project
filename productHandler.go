package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

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
