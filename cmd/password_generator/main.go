package main

import (
	"log"

	v1router "github.com/christoffer1009/password-generator/api/v1/router"
	v2router "github.com/christoffer1009/password-generator/api/v2/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	v1router.SetupRoutesV1(app)
	v2router.SetupRoutesV2(app)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}
