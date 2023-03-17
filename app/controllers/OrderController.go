package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/lambda-platform/lambda/DB"
	agentUtils "github.com/lambda-platform/lambda/agent/utils"
	"lambda/app/models"
	"net/http"
)

func CreateOrder(c *fiber.Ctx) error {
	orderUser := agentUtils.AuthUserObject(c)

	var cartZahialga []models.ViewCartZahialga
	DB.DB.Where("user_id = ?", orderUser["id"]).Order("id DESC").Find(&cartZahialga)

	var totalPrice = 0

	for _, cartPrice := range cartZahialga {
		fmt.Println(cartPrice.Price)

		totalPrice = totalPrice + cartPrice.Price

	}
	fmt.Println("total", totalPrice)

	newOrder := new(models.Orders)

	newOrder.UserID = GetIntegerPointer(int(orderUser["id"].(int64)))
	newOrder.PaymentType = GetStringPointer("")
	newOrder.Price = totalPrice

	DB.DB.Create(&newOrder)

	for _, cartDetail := range cartZahialga {
		orderDetail := models.OrderDetail{}

		orderDetail.OrderID = newOrder.ID
		orderDetail.UserID = GetIntegerPointer(int(orderUser["id"].(int64)))
		orderDetail.FoodID = cartDetail.FoodID
		orderDetail.CartID = cartDetail.ID
		orderDetail.Price = cartDetail.Price

		DB.DB.Create(&orderDetail)

	}

	return c.JSON(newOrder)

}

func CancelOrder(c *fiber.Ctx) error {
	order := new(models.OrdersStatus)

	orderCancelReq := models.OrderRequest{}
	errReq := c.BodyParser(&orderCancelReq)
	if errReq != nil {
		fmt.Println(errReq.Error())
		return c.Status(http.StatusInternalServerError).JSON("server error")
	}

	fmt.Println(orderCancelReq.ID)

	orderUser := agentUtils.AuthUserObject(c)

	order.PaymentStatus = "canceled"
	DB.DB.Where("id = ? AND user_id = ? AND payment_status = 'pending'", orderCancelReq.ID, orderUser["id"]).Order("id DESC").Find(&order)

	if order.ID == 0 {
		return c.Status(http.StatusOK).JSON(map[string]interface{}{
			"status":  "warning",
			"message": "Not found order",
		})
	}

	order.PaymentStatus = "canceled"

	DB.DB.Save(&order)

	return c.Status(http.StatusOK).JSON(map[string]interface{}{
		"status":  "success",
		"message": "Захиалга цуцлгадлаа",
	})
}

func GetIntegerPointer(value int) int {
	return value
}

func GetStringPointer(value string) string {
	return value
}
