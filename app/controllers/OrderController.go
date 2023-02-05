package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/lambda-platform/lambda/DB"
	"lambda/app/models"
	"net/http"
)

func CreateOrder(c *fiber.Ctx) error {

	order := &models.Orders{}
	err := c.BodyParser(order)

	var orderView []models.OrdersView
	DB.DB.Table("orders").Where("user_id = ?", order.UserID).Find(&orderView)
	if len(orderView) > 0 {
		return c.Status(http.StatusOK).JSON(map[string]string{
			"status":  "warning",
			"message": "Захиалга үүсгэсэн байна.",
		})
	}

	if err != nil {
		fmt.Println(err.Error())
		return c.Status(http.StatusInternalServerError).JSON("server error")
	}

	//if orderView.UserID >= 1 {
	DB.DB.Create(&order)
	//}

	return c.Status(http.StatusOK).JSON(map[string]string{
		"status":  "success",
		"message": "Захиалга үүсгэгдсэн",
	})
}

func GetOrders(c *fiber.Ctx) error {
	orderView := []models.OrdersView{}
	DB.DB.Find(&orderView)

	return c.JSON(orderView)
}
