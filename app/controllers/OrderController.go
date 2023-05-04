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

func GetPendingOrder(c *fiber.Ctx) error {
	activeOrder := new(models.MyOrders)
	MyOrder := new(models.MyOrders)
	orderUser := agentUtils.AuthUserObject(c)

	DB.DB.Select("id, order_number, created_at, ebarimt_data, ebarimt_type, org_register_number, payment_at, payment_status, payment_type, quantity, total_price, user_id, EXTRACT(EPOCH FROM (now() - created_at)) as uldsen_sec").Table("orders").Where("age(now(), created_at) < '5 minutes' AND user_id = ? AND payment_status = ?").Order("id DESC").First(&activeOrder)
	DB.DB.Where("user_id = ? AND payment_status = ?", orderUser["id"], "pending").Order("id DESC").First(&MyOrder)

	if activeOrder.ID >= 1 {
		return c.Status(http.StatusOK).JSON(map[string]interface{}{
			"status": "success",
			"order":  activeOrder,
		})
	} else {
		return c.Status(http.StatusOK).JSON(map[string]interface{}{
			"status":       "warning",
			"message":      "Худалдан авалтын хугцаа дууссан байна",
			"order_number": MyOrder.OrderNumber,
		})
	}
}

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
		orderType := ""

		for _, cartZahialgaQty := range cartZahialga {
			totalQtyZahialga = totalQtyZahialga + cartZahialgaQty.Qty
			totalPriceZahialge = totalPriceZahialge + cartZahialgaQty.Price
			cartID = cartZahialgaQty.ID
			orderType = "Захиалгат"
		}

		for _, cartMenu := range cartMenus {
			totalQtyMenu = totalQtyMenu + cartMenu.Qty
			totalPriceMenu = totalPriceMenu + cartMenu.PacketPrice
			orderType = "Багц"
		}

		totalQty := totalQtyMenu + totalQtyZahialga
		totalPrice := totalPriceMenu + totalPriceZahialge

		orders.UserID = GetIntegerPointer(int(orderUser["id"].(int64)))
		orders.PaymentStatus = GetStringPointer("pending")
		orders.OrderQuantity = totalQty
		orders.Price = totalPrice
		orders.IsSelled = GetStringPointer("olgoogui")
		orders.CartID = cartID
		orders.OrderType = orderType
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

	orderCancelReq := models.Orders{}
	errReq := c.BodyParser(&orderCancelReq)
	if errReq != nil {
		fmt.Println(errReq.Error())
		return c.Status(http.StatusInternalServerError).JSON("server error")
	}

	orderUser := agentUtils.AuthUserObject(c)

	//DB.DB.Model(models.OrdersStatus{}).Where("id = ? AND user_id = ? AND payment_status = 'pending'", orderCancelReq.ID, orderUser["id"]).Order("id DESC").Update("payment_status", "canceled")

	DB.DB.Where("id = ? AND user_id = ? AND payment_status = 'pending'", orderCancelReq.ID, orderUser["id"]).Find(&orderCancelReq)

	DB.DB.Delete(&orderCancelReq)

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

//func ChekingBalance(c *fiber.Ctx, oldOrders models.Orders, order models.ViewOrder) {
//	orderUser := agentUtils.AuthUserObject(c)
//
//	var orderDetails []models.OrderDetail
//	DB.DB.Where("order_id = ? AND user_id = ?", oldOrders.ID, orderUser["id"]).Find(&orderDetails)
//
//	var orderDetailSets []models.OrderDetailSet
//	DB.DB.Where("order_id = ?", oldOrders.ID).Find(&orderDetailSets)
//
//	for _, orderDetailSet := range orderDetailSets {
//		foodBalance := models.ViewFoodBalance{}
//		DB.DB.Where("food_id = ? AND kitchen_id = ?", orderDetailSet.FoodID, orderDetailSet.KitchenID).Find(&foodBalance)
//
//		if foodBalance.Quantity < orderDetailSet.Quantity {
//			return nil
//			}
//		}
//		UpdateStatus(order.OrderNumber, oldOrders.ID, "mmk", "success")
//		UpdateBalance(orderDetailSet.FoodID, orderDetailSet.KitchenID, orderDetailSet.Quantity)
//	}
//
//	for _, orderDetail := range orderDetails {
//		UpdateBalance(orderDetail.FoodID, orderDetail.KitchenID, orderDetail.Qty)
//	}
//}

func UpdateStatus(OrderNumber string, OrderID int, PaymentType string, PaymentStatus string) {
	now := time.Now()
	editOrder := models.Orders{}

	DB.DB.Debug().Model(&editOrder).Where("id = ? AND order_number = ?", OrderID, OrderNumber).Updates(models.Orders{PaymentStatus: PaymentStatus, PaymentType: PaymentType, SuccessTime: now.Format("2006-01-02 15:04:05")})

}

func UpdateBalance(FoodID int, KitchenID int, Quantity int) {

	foodBalance := models.FoodBalance{}
	DB.DB.Where("food_id = ? AND kitchen_id = ?", FoodID, KitchenID).Find(&foodBalance)

	//if foodBalance.Quantity < Quantity {
	//	c.Status(http.StatusOK).JSON(models.UpdateResponse{
	//		Status:  "warning",
	//		Message: "Сонгосон хоолны үлдэгдэл хүрэлцэхгүй байна",
	//	})

	//resp := models.UpdateResponse{
	//	Status:  "warning",
	//	Message: "Сонгосон хоолны үлдэгдэл хүрэлцэхгүй байна",
	//}
	//c.Status(http.StatusOK).JSON(resp)
	//return
	//}

	balanceQty := foodBalance.Quantity - Quantity

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
