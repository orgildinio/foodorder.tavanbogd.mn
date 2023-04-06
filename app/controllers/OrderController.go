package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/lambda-platform/lambda/DB"
	agentUtils "github.com/lambda-platform/lambda/agent/utils"
	"lambda/app/models"
	"log"
	"net/http"
	"time"
)

func CreateOrder(c *fiber.Ctx) error {
	orderUser := agentUtils.AuthUserObject(c)
	orders := models.Orders{}

	var cartMenus []models.CartMenu
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

		cartID := 0

		for _, cartZahialgaQty := range cartZahialga {
			totalQtyZahialga = totalQtyZahialga + cartZahialgaQty.Qty
			totalPriceZahialge = totalPriceZahialge + cartZahialgaQty.Price
			cartID = cartZahialgaQty.ID

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
		orders.CartID = cartID
		DB.DB.Omit("is_delivery").Create(&orders)

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
				orderDetail.OrderType = GetStringPointer("zahialgat")

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
						orderDetailSet := models.OrderDetailSet{}

						orderDetailSet.UserID = GetIntegerPointer(int(orderUser["id"].(int64)))
						orderDetailSet.OrderID = orders.ID
						orderDetailSet.FoodID = cartSubMenuFood.FoodID
						orderDetailSet.CartID = cartMenu.ID
						orderDetailSet.KitchenID = cartMenu.KitchenID
						orderDetailSet.Quantity = GetIntegerPointer(1)

						DB.DB.Create(&orderDetailSet)

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
				orderDetail.OrderType = GetStringPointer("bagts")

				DB.DB.Create(&orderDetail)
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
	now := time.Now()
	editOrder := models.Orders{}

	DB.DB.Debug().Model(&editOrder).Where("id = ? AND order_number = ?", OrderID, OrderNumber).Updates(models.Orders{PaymentStatus: PaymentStatus, PaymentType: PaymentType, SuccessTime: now.Format("2006/01/02 15:04:05")})

}

func UpdateBalance(FoodID int, KitchenID int, Quantity int) {

	foodBalance := models.FoodBalance{}
	DB.DB.Where("food_id = ? AND kitchen_id = ?", FoodID, KitchenID).Find(&foodBalance)

	balanceQty := foodBalance.Quantity - Quantity

	log.Println("balanceQty", balanceQty)

	DB.DB.Model(&foodBalance).Where("food_id = ? AND kitchen_id = ?", FoodID, KitchenID).Update("quantity", balanceQty)

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

	DB.DB.Debug().Omit("is_delivery").Save(&order)

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
