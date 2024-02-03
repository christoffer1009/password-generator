package router

import (
	"github.com/christoffer1009/password-generator/api/v1/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutesV1(app *fiber.App) {
	app.Post("/v1/generate", handler.Generate)
}
