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

	cartMenus := []models.CartMenu{}
	DB.DB.Where("user_id = ?", orderUser["id"]).Order("id DESC").Find(&cartMenus)

	var cartZahialga []models.ViewCartZahialga
	DB.DB.Where("user_id = ?", orderUser["id"]).Order("id DESC").Find(&cartZahialga)

	var viewOrder []models.ViewOrder
	DB.DB.Where("user_id = ? AND payment_status = 'pending'", orderUser["id"]).Find(&viewOrder)

	packetPrice := models.LutPacketPrice{}
	DB.DB.Find(&packetPrice)

	if len(viewOrder) >= 1 {
		return c.Status(http.StatusOK).JSON(map[string]interface{}{
			"status":  "warning",
			"message": "Таньд идэвхтэй захиалга байна",
		})
	}

	if len(cartMenus) >= 1 || len(cartZahialga) >= 1 {

		totalQtyMenu := 0
		totalQtyZahialga := 0

		totalPriceMenu := 0
		totalPriceZahialge := 0

		for _, cartZahialgaQty := range cartZahialga {
			totalQtyZahialga = totalQtyZahialga + cartZahialgaQty.Qty
			totalPriceZahialge = totalPriceZahialge + cartZahialgaQty.Price
		}

		for _, cartMenu := range cartMenus {
			totalQtyMenu = totalQtyMenu + cartMenu.Qty
			totalPriceMenu = totalPriceMenu + cartMenu.PacketPrice
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
				foodBalance := models.FoodBalance{}
				DB.DB.Where("food_id = ? AND kitchen_id = ?", cartZahialgas.FoodID, cartZahialgas.KitchenID).Find(&foodBalance)

				orderDetail := models.OrderDetail{}

				orderDetail.UserID = GetIntegerPointer(int(orderUser["id"].(int64)))
				orderDetail.OrderID = orders.ID
				orderDetail.FoodID = cartZahialgas.FoodID
				orderDetail.KitchenID = cartZahialgas.KitchenID
				orderDetail.CartID = cartZahialgas.ID
				orderDetail.Qty = GetIntegerPointer(1)
				orderDetail.Price = int(foodBalance.FoodPrice)

				if foodBalance.Quantity < orderDetail.Qty {
					return c.Status(http.StatusOK).JSON(map[string]string{
						"status":    "warning",
						"status_mn": "Анхааруулга",
						"message":   "Сонгосон хоолны үлдэгдэл хүрэлцэхгүй байна",
					})
				}

				DB.DB.Create(&orderDetail)

			}

			zahialgatData := models.CartZahialgat{}
			DB.DB.Where("user_id = ?", orderUser["id"]).Order("id DESC").Find(&zahialgatData)
			DB.DB.Delete(zahialgatData)
		}

		for _, cartMenu := range cartMenus {

			var cartSubMenus []models.CartSubMenu
			DB.DB.Where("menu_id = ?", cartMenu.ID).Find(&cartSubMenus)

			for _, cartSubMenu := range cartSubMenus {
				var cartSubMenuFoods []models.CartSubMenuFood
				DB.DB.Where("sub_menu_id = ?", cartSubMenu.ID).Find(&cartSubMenuFoods)

				for _, cartSubMenuFood := range cartSubMenuFoods {
					for i := 1; i <= cartMenu.Qty; i++ {

						foodBalance := models.FoodBalance{}
						DB.DB.Where("food_id = ? AND kitchen_id = ?", cartSubMenuFood.FoodID, cartMenu.KitchenID).Find(&foodBalance)

					}
				}
			}
			for i := 1; i <= cartMenu.Qty; i++ {
				orderDetail := models.OrderDetail{}

				orderDetail.UserID = GetIntegerPointer(int(orderUser["id"].(int64)))
				orderDetail.OrderID = orders.ID
				orderDetail.MenuID = cartMenu.MenuID
				orderDetail.Qty = GetIntegerPointer(1)
				orderDetail.Price = int(packetPrice.PacketPrice)
				orderDetail.KitchenID = cartMenu.KitchenID
				orderDetail.CartID = cartMenu.CartID

				DB.DB.Save(&orderDetail)
			}

		}

		cartBagts := models.CartMenu{}
		DB.DB.Debug().Where("user_id = ?", orderUser["id"]).Order("id DESC").Delete(&cartBagts)

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

	orderCancelReq := models.OrderRequest{}
	errReq := c.BodyParser(&orderCancelReq)
	if errReq != nil {
		fmt.Println(errReq.Error())
		return c.Status(http.StatusInternalServerError).JSON("server error")
	}

	orderUser := agentUtils.AuthUserObject(c)

	DB.DB.Model(models.OrdersStatus{}).Where("id = ? AND user_id = ? AND payment_status = 'pending'", orderCancelReq.ID, orderUser["id"]).Order("id DESC").Update("payment_status", "canceled")

	return c.Status(http.StatusOK).JSON(map[string]interface{}{
		"status":  "success",
		"message": " Захиалга цуцлгадлаа",
	})
}

func DeleteOrder() {
	var orders []models.Orders
	DB.DB.Debug().Where("age(created_at) < '15 minute' AND payment_status = 'pending'").Find(&orders)

	for _, order := range orders {
		DB.DB.Debug().Delete(&order)
	}

}

func UpdateStatus(OrderNumber string, OrderID int, PaymentType string, PaymentStatus string) {

	editOrder := models.Orders{}
	DB.DB.Debug().Where("id = ? AND order_number = ?", OrderID, OrderNumber).Find(&editOrder)

	now := time.Now()

	editOrder.PaymentStatus = PaymentStatus
	editOrder.PaymentType = PaymentType
	editOrder.SuccessTime = now.Format("2006-02-01 15:04:05")

	//DB.DB.Debug().Save(&editOrder)

}

func UpdateBalance(FoodID int, KitchenID int, Quantity int) {
	foodBalance := models.FoodBalance{}
	DB.DB.Where("food_id = ? AND kitchen_id = ?", FoodID, KitchenID).Find(&foodBalance)

	foodBalance.Quantity = foodBalance.Quantity - Quantity

	DB.DB.Save(foodBalance)

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
