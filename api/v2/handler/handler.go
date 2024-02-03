package handler

import (
	"github.com/christoffer1009/password-generator/internal/entity"
	"github.com/christoffer1009/password-generator/internal/pkg/generator"
	"github.com/gofiber/fiber/v2"
)

func Generate(c *fiber.Ctx) error {
	var opt entity.Options
	opt.Length = c.QueryInt("length", 6)
	opt.UpperCase = c.QueryBool("upperCase", false)
	opt.LowerCase = c.QueryBool("lowerCase", true)
	opt.Digits = c.QueryBool("digits", false)
	opt.SpecialChars = c.QueryBool("specialChars", false)

	password, err := generator.GeneratePassword(&opt)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{"password": password})
}
