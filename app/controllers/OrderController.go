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

	var checkOldOrders []models.Orders
	DB.DB.Where("user_id = ? AND payment_status = 'pending'", orderUser["id"]).Order("id DESC").Find(&checkOldOrders)

	for _, checkOldOrder := range checkOldOrders {
		checkOldOrder.PaymentStatus = "canceled"
		DB.DB.Save(&checkOldOrder)
	}

	var totalPrice = 0
	var totalQty = 0

	for _, cartPrice := range cartZahialga {
		fmt.Println(cartPrice.Price)

		totalPrice = totalPrice + cartPrice.Price
		totalQty = totalQty + cartPrice.Qty

	}

	fmt.Println(totalQty)

	newOrder := new(models.Orders)

	newOrder.UserID = GetIntegerPointer(int(orderUser["id"].(int64)))
	newOrder.PaymentType = GetStringPointer("")
	newOrder.Price = totalPrice
	newOrder.OrderQuantity = totalQty
	newOrder.OrderType = GetStringPointer("zahialgat")

	DB.DB.Create(&newOrder)

	for _, cartDetail := range cartZahialga {
		orderDetail := models.OrderDetail{}

		orderDetail.OrderID = newOrder.ID
		orderDetail.UserID = GetIntegerPointer(int(orderUser["id"].(int64)))
		orderDetail.FoodID = cartDetail.FoodID
		orderDetail.CartID = cartDetail.ID
		orderDetail.Price = cartDetail.Price
		orderDetail.Qty = cartDetail.Qty

		DB.DB.Create(&orderDetail)

		DB.DB.Where("id = ?", cartDetail.ID).Delete(&cartDetail)

	}

	return c.Status(http.StatusOK).JSON(map[string]interface{}{
		"status":  "success",
		"message": "Захиалга үүсгэлээ",
	})

}

func CancelOrder(c *fiber.Ctx) error {
	order := new(models.OrdersStatus)

	orderCancelReq := models.OrderRequest{}
	errReq := c.BodyParser(&orderCancelReq)
	if errReq != nil {
		fmt.Println(errReq.Error())
		return c.Status(http.StatusInternalServerError).JSON("server error")
	}

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
		"message": order.OrderNumber + " дугаартай захиалга цуцлгадлаа",
	})
}

func CreateOrderSet(c *fiber.Ctx) error {
	orderUser := agentUtils.AuthUserObject(c)

	var viewCartMenu []models.ViewCartMenu
	DB.DB.Where("user_id = ?", orderUser["id"]).Order("id DESC").Find(&viewCartMenu)

	if len(viewCartMenu) == 0 {
		return c.Status(http.StatusOK).JSON(map[string]interface{}{
			"status":  "warning",
			"message": "User not found",
		})
	}
	totalQty := 0
	var totalPrice float32

	cartID := 0

	for _, cartDetail := range viewCartMenu {

		totalQty = totalQty + cartDetail.Qty
		totalPrice = totalPrice + cartDetail.PacketPrice
		cartID = cartDetail.ID

	}

	newOrder := new(models.Orders)

	newOrder.UserID = GetIntegerPointer(int(orderUser["id"].(int64)))
	newOrder.PaymentType = GetStringPointer("")
	newOrder.Price = int(totalPrice)
	newOrder.OrderQuantity = totalQty
	newOrder.OrderType = GetStringPointer("bagts")

	DB.DB.Create(&newOrder)

	cartMenu := models.CartMenu{}

	fmt.Println("Cart ID", cartID)

	DB.DB.Where("id = ? AND user_id = ?", cartID, orderUser["id"]).Delete(&cartMenu)

	return c.Status(http.StatusOK).JSON(map[string]interface{}{
		"status":  "success",
		"message": "Захиалга үүсгэлээ",
	})
}

func GetIntegerPointer(value int) int {
	return value
}

func GetStringPointer(value string) string {
	return value
}
