package controllers

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"os"
)

func Home(c *fiber.Ctx) error {
	return c.Status(http.StatusSeeOther).Redirect("/auth/login")
}
func Admin(c *fiber.Ctx) error {

	return JustRenderTemplate("public/admin/index.html", c)
}
func Login(c *fiber.Ctx) error {

	return JustRenderTemplate("public/auth/index.html", c)
}
func JustRenderTemplate(file string, c *fiber.Ctx) error {
	buff, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	c.Set("Content-Type", "text/html")
	return c.SendString(string(buff))
}
