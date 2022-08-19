package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func HomeProduction(c *fiber.Ctx) error {
	return c.Render("public/index", map[string]interface{}{})
}
