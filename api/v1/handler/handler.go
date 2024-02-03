package handler

import (
	"encoding/json"

	"github.com/christoffer1009/password-generator/internal/entity"
	"github.com/christoffer1009/password-generator/internal/pkg/generator"
	"github.com/gofiber/fiber/v2"
)

func Generate(c *fiber.Ctx) error {
	var opt entity.Options
	if err := json.Unmarshal([]byte(c.Body()), &opt); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error decoding JSON")
	}

	password, err := generator.GeneratePassword(&opt)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{"password": password})
}
