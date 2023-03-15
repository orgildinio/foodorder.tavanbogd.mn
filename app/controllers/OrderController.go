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
	order := new(models.Orders)
	err := c.BodyParser(&order)

	if err != nil {
		fmt.Println(err.Error())
		return c.Status(http.StatusInternalServerError).JSON("server error")
	}

	orderUser := agentUtils.AuthUserObject(c)
	order.UserID = int(orderUser["id"].(int64))

	orderCheck := models.CheckOrders{}
	DB.DB.Debug().Where("user_id = ? AND cart_id = ? AND age(created_at) <= '15 minutes' AND order_status = 'pending'", order.UserID, order.CartID).Order("id DESC").Find(&orderCheck)

	if orderCheck.ID >= 1 {
		return c.JSON(map[string]string{
			"status":      "warning",
			"orderNumber": "Таньд " + *orderCheck.OrderNumber + " дугаартай захиалга үүссэн байна.",
			"status_mn":   "Анхааруулга",
		})
	}

	cartZahialga := models.CartZahialgat{}
	DB.DB.Debug().Where("id = ?", order.CartID).Find(&cartZahialga)

	fmt.Println("zahialsan_id", cartZahialga.FoodID)

	foodBalance := models.FoodBalance{}
	DB.DB.Where("food_id = ?", cartZahialga.FoodID).Find(&foodBalance)

	fmt.Println("balansiin_id", foodBalance.FoodID)

	order.OrderStatus = GetStringPointer("pending")
	//foodBalance.Qty = foodBalance.Qty - cartZahialga.Qty

	DB.DB.Debug().Save(&foodBalance)

	//zahialgatStr := 'zahialgat'
	//
	//if order.CartType == 'zahialgat' {
	//}

	DB.DB.Create(&order)

	return c.Status(http.StatusOK).JSON(map[string]interface{}{
		"status":  "success",
		"message": "Захиалга үүсгэлээ.",
	})
}

func CancelOrder(c *fiber.Ctx) error {
	order := new(models.Orders)
	err := c.BodyParser(&order)

	if err != nil {
		fmt.Println(err.Error())
		return c.Status(http.StatusInternalServerError).JSON("server error")
	}

	orderUser := agentUtils.AuthUserObject(c)
	order.UserID = int(orderUser["id"].(int64))

	checkOrder := models.ViewOrders{}
	DB.DB.Debug().Where("id = ?", order.ID).Order("id DESC").Find(&checkOrder)

	order.OrderStatus = GetStringPointer("cancel")

	DB.DB.Save(&order)

	return c.Status(http.StatusOK).JSON(map[string]interface{}{
		"status":  "success",
		"message": checkOrder.OrderNumber + " дугаартай захиалга цуцлагдлаа",
	})
}

func GetStringPointer(value string) *string {
	return &value
}
