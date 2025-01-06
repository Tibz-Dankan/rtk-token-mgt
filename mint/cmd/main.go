package main

import (
	"fmt"
	"log"

	"github.com/Tibz-Dankan/rtk-token-mgt/internal/handlers"
	"github.com/Tibz-Dankan/rtk-token-mgt/internal/pkg"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: pkg.ErrorHandler,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Use(logger.New())

	mint := app.Group("/api/v0/mint", func(c *fiber.Ctx) error {
		return c.Next()
	})
	mint.Post("/token/new", handlers.CreateToken)

	app.Use("*", func(c *fiber.Ctx) error {
		message := fmt.Sprintf("api route '%s' doesn't exist!", c.Path())
		return fiber.NewError(fiber.StatusNotFound, message)
	})

	log.Fatal(app.Listen(":5000"))
}
