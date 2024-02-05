package main

import (
	"log"

	v1router "github.com/christoffer1009/password-generator/api/v1/router"
	v2router "github.com/christoffer1009/password-generator/api/v2/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Set("Access-Control-Allow-Headers", "Content-Type")

		return c.Next()
	})

	v1router.SetupRoutesV1(app)
	v2router.SetupRoutesV2(app)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}
