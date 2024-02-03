package router

import (
	"github.com/christoffer1009/password-generator/api/v2/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutesV2(app *fiber.App) {
	app.Get("/v2/generate", handler.Generate)
}
