package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/lambda-platform/lambda/DB"
	agentUtils "github.com/lambda-platform/lambda/agent/utils"
	"lambda/app/models"
	"net/http"
	"time"
)

func CreateOrder(c *fiber.Ctx) error {
	orderUser := agentUtils.AuthUserObject(c)

	orders := models.Orders{}

	var cartZahialga []models.ViewCartZahialga
	DB.DB.Where("user_id = ?", orderUser["id"]).Order("id DESC").Find(&cartZahialga)

	var cartMenu []models.ViewCartMenu
	DB.DB.Where("user_id = ?", orderUser["id"]).Order("id DESC").Find(&cartMenu)

	viewOrder := []models.ViewOrder{}
	DB.DB.Where("user_id = ? AND payment_status = 'pending'", orderUser["id"]).Find(&viewOrder)

	//if len(cartZahialga) == 0 || len(cartMenu) == 0 {
	//
	//	return c.Status(http.StatusOK).JSON(map[string]interface{}{
	//		"status":  "warning",
	//		"message": "Not Found Carting Items",
	//	})
	//}

	if len(viewOrder) >= 1 {

		return c.Status(http.StatusOK).JSON(map[string]interface{}{
			"status":  "warning",
			"message": "Таньд идэвхтэй захиалга байна",
		})
	}

	if len(cartMenu) >= 1 || len(cartZahialga) >= 1 {

		totalQtyMenu := 0
		totalQtyZahialga := 0

		totalPriceMenu := 0
		totalPriceZahialge := 0

		for _, cartMenuQty := range cartMenu {
			totalQtyMenu = totalQtyMenu + cartMenuQty.Qty
			totalPriceMenu = totalPriceMenu + int(cartMenuQty.PacketPrice)
		}

		for _, cartZahialgaQty := range cartZahialga {
			totalQtyZahialga = totalQtyZahialga + cartZahialgaQty.Qty
			totalPriceZahialge = totalPriceZahialge + cartZahialgaQty.Price

			fmt.Println(cartZahialgaQty.Price)
		}

		totalQty := totalQtyMenu + totalQtyZahialga
		totalPrice := totalPriceMenu + totalPriceZahialge

		orders.UserID = GetIntegerPointer(int(orderUser["id"].(int64)))
		orders.PaymentStatus = GetStringPointer("pending")
		orders.OrderQuantity = totalQty
		orders.Price = totalPrice
		orders.IsSelled = GetStringPointer("olgoogui")

		DB.DB.Create(&orders)

		for _, cartZahialgas := range cartZahialga {
			for i := 1; i <= cartZahialgas.Qty; i++ {
				orderDetail := models.OrderDetail{}

				orderDetail.UserID = GetIntegerPointer(int(orderUser["id"].(int64)))
				orderDetail.OrderID = orders.ID
				orderDetail.FoodID = cartZahialgas.FoodID
				orderDetail.KitchenID = cartZahialgas.KitchenID
				orderDetail.CartID = cartZahialgas.ID
				orderDetail.Qty = cartZahialgas.Qty
				orderDetail.Price = cartZahialgas.Price

				DB.DB.Create(&orderDetail)
			}

			zahialgatData := models.CartZahialgat{}
			DB.DB.Where("user_id = ?", orderUser["id"]).Order("id DESC").Find(&zahialgatData)
			DB.DB.Delete(zahialgatData)
		}

		for _, cartMenus := range cartMenu {
			//for j := 1; j <= cartMenus.Qty; j++ {
			orderDetail := models.OrderDetail{}

			orderDetail.UserID = GetIntegerPointer(int(orderUser["id"].(int64)))
			orderDetail.OrderID = orders.ID
			orderDetail.MenuID = cartMenus.MenuID
			orderDetail.KitchenID = cartMenus.KitchenID
			orderDetail.Qty = cartMenus.Qty
			orderDetail.Price = int(cartMenus.PacketPrice)
			orderDetail.CartID = cartMenus.ID

			DB.DB.Create(&orderDetail)
			//}

			cartBagts := models.CartMenu{}
			DB.DB.Where("user_id = ?", orderUser["id"]).Order("id DESC").Find(&cartBagts)
			DB.DB.Delete(&cartBagts)
		}

	}

	ordersResponse := models.ViewOrder{}
	DB.DB.Where("user_id = ?", orderUser["id"]).Order("id DESC").Find(&ordersResponse)

	return c.Status(http.StatusOK).JSON(map[string]interface{}{
		"status":  "success",
		"message": "Захиалга үүсгэлээ",
		"data":    ordersResponse,
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

	DB.DB.Where("id = ? AND user_id = ? AND payment_status = 'pending'", orderCancelReq.ID, orderUser["id"]).Order("id DESC").Find(&order)

	if order.ID == 0 {
		return c.Status(http.StatusOK).JSON(map[string]interface{}{
			"status":  "warning",
			"message": "Not found order",
		})
	}

	DB.DB.Delete(&order)

	return c.Status(http.StatusOK).JSON(map[string]interface{}{
		"status":  "success",
		"message": order.OrderNumber + " дугаартай захиалга цуцлгадлаа",
	})
}

func UpdateStatus(UserID interface{}, OrderID int, PaymentType string, PaymentStatus string) {

	editOrder := models.Orders{}
	DB.DB.Where("id = ? AND user_id = ?", OrderID, UserID).Find(&editOrder)

	if OrderID > 0 {

		now := time.Now()

		editOrder.PaymentStatus = PaymentStatus
		editOrder.PaymentType = PaymentType
		editOrder.SuccessTime = now.Format("2006-02-01 15:04:05")

		DB.DB.Save(&editOrder)

		UpdateBalance(UserID, OrderID)
	} else {
		panic("Not found order")
	}

}

func UpdateBalance(UserID interface{}, OrderID int) {

	var orderDetails []models.OrderDetail
	DB.DB.Where("user_id = ? AND order_id = ?", UserID, OrderID).Find(&orderDetails)

	if OrderID > 0 {

		for _, items := range orderDetails {
			foodBalance := models.FoodBalance{}
			DB.DB.Where("food_id = ? AND kitchen_id = ?", items.FoodID, items.KitchenID).Find(&foodBalance)

			foodBalance.Quantity = foodBalance.Quantity - items.Qty

			DB.DB.Save(&foodBalance)
		}
	} else {
		panic("Not found Order ID")
	}

}

func RecepcionRequest(c *fiber.Ctx) error {
	receptionRequestData := models.ReceptionRequestData{}
	err := c.BodyParser(&receptionRequestData)
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(http.StatusInternalServerError).JSON("server error")
	}

	order := models.Orders{}
	DB.DB.Where("id = ? AND is_selled = 'olgoogui'", receptionRequestData.ID).Find(&order)

	if order.ID == 0 {
		return c.Status(http.StatusOK).JSON(map[string]interface{}{
			"status":  "warning",
			"message": "Order not found",
		})
	}

	order.IsSelled = receptionRequestData.PaymentStatus

	DB.DB.Save(&order)

	return c.Status(http.StatusOK).JSON(map[string]interface{}{
		"status":  "success",
		"message": "Хоол хүлээлгэн өгсөн",
	})
}

func GetIntegerPointer(value int) int {
	return value
}

func GetStringPointer(value string) string {
	return value
}
