package controllers

import (
	"github.com/gofiber/fiber/v2"
)

//home page
func HomeProduction(c *fiber.Ctx) error {
	return c.Render("home", map[string]interface{}{})
}

func HomeData() map[string]interface{} {

	return map[string]interface{}{
		"title": "test",
	}
}
