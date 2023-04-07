package controllers

import (
	"github.com/gofiber/fiber/v2"
	"lambda/ebarimt"
)

func EBarimtInfo(c *fiber.Ctx) error {

	info, err := ebarimt.PosAPI.GetInformation()
	if err != nil {
		return err
	}

	return c.JSON(info)
}

func EBarimtCheck(c *fiber.Ctx) error {

	res, err := ebarimt.PosAPI.CheckApi()
	if err != nil {
		return err
	}

	return c.JSON(res)
}

func EBarimtSend(c *fiber.Ctx) error {

	res, err := ebarimt.PosAPI.SendData()
	if err != nil {
		return err
	}

	return c.JSON(res)
}
