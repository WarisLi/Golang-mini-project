package middleware

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func AppLogger(c *fiber.Ctx) error {
	start := time.Now()
	fmt.Printf("URL = %s, Method = %s, Time = %s\n", c.OriginalURL(), c.Method(), start)

	return c.Next()
}
