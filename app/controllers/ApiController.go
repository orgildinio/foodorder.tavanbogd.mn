package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lambda-platform/lambda/DB"
	"lambda/app/models"
)

func Users(c *fiber.Ctx) error {

	users := []models.Users{}
	DB.DB.Find(&users)

	return c.JSON(users)
}
