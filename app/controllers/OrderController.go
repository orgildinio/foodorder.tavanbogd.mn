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

	orders := models.Orders{}

	var cartZahialga []models.ViewCartZahialga
	DB.DB.Where("user_id = ?", orderUser["id"]).Order("id DESC").Find(&cartZahialga)

	var cartMenu []models.ViewCartMenu
	DB.DB.Where("user_id = ?", orderUser["id"]).Order("id DESC").Find(&cartMenu)

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
		}

		totalQty := totalQtyMenu + totalQtyZahialga
		totalPrice := totalPriceMenu + totalPriceZahialge

		orders.UserID = GetIntegerPointer(int(orderUser["id"].(int64)))
		orders.PaymentStatus = GetStringPointer("pending")
		orders.OrderQuantity = totalQty
		orders.Price = totalPrice

		DB.DB.Create(&orders)

		for _, cartZahialgas := range cartZahialga {
			//for i := 1; i <= cartZahialgas.Qty; i++ {
			orderDetail := models.OrderDetail{}

			orderDetail.UserID = GetIntegerPointer(int(orderUser["id"].(int64)))
			orderDetail.OrderID = orders.ID
			orderDetail.FoodID = cartZahialgas.FoodID
			orderDetail.KitchenID = cartZahialgas.KitchenID
			orderDetail.CartID = cartZahialgas.ID
			orderDetail.Qty = cartZahialgas.Qty
			orderDetail.Price = cartZahialgas.Price

			DB.DB.Create(&orderDetail)
			//}

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
			DB.DB.Delete(cartBagts)
		}

	}

	return c.Status(http.StatusOK).JSON(map[string]string{
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

//func CreateOrderSet(c *fiber.Ctx) error {
//    orderUser := agentUtils.AuthUserObject(c)
//
//    var viewCartMenu []models.ViewCartMenu
//    DB.DB.Where("user_id = ?", orderUser["id"]).Order("id DESC").Find(&viewCartMenu)
//
//    if len(viewCartMenu) == 0 {
//        return c.Status(http.StatusOK).JSON(map[string]interface{}{
//            "status":  "warning",
//            "message": "User not found",
//        })
//    }
//    totalQty := 0
//    var totalPrice float32
//
//    for _, cartDetail := range viewCartMenu {
//
//        totalQty = totalQty + cartDetail.Qty
//        totalPrice = float32(totalQty) * cartDetail.PacketPrice
//    }
//
//    newOrder := new(models.Orders)
//
//    newOrder.UserID = GetIntegerPointer(int(orderUser["id"].(int64)))
//    newOrder.PaymentType = GetStringPointer("")
//    newOrder.Price = int(totalPrice)
//    newOrder.OrderQuantity = totalQty
//    newOrder.OrderType = GetStringPointer("bagts")
//    newOrder.PaymentStatus = GetStringPointer("pending")
//
//    DB.DB.Create(&newOrder)
//
//    for _, viewCartMenus := range viewCartMenu {
//        orderDetail := models.OrderDetail{}
//
//        orderDetail.OrderID = newOrder.ID
//        orderDetail.UserID = GetIntegerPointer(int(orderUser["id"].(int64)))
//        orderDetail.CartID = viewCartMenus.ID
//    }
//
//    cartMenu := models.CartMenu{}
//    DB.DB.Where("user_id = ?", orderUser["id"]).Order("id DESC").Find(&cartMenu)
//    DB.DB.Delete(&cartMenu)
//
//    return c.Status(http.StatusOK).JSON(map[string]interface{}{
//        "status":  "success",
//        "message": "Захиалга үүсгэлээ",
//    })
//}

func GetIntegerPointer(value int) int {
	return value
}

func GetStringPointer(value string) string {
	return value
}
