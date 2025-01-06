package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

var CreateToken = func(c *fiber.Ctx) error {

	fmt.Println("Trigger creating token here")

	response := map[string]interface{}{
		"status":  "success",
		"message": "Token created successfully",
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}
