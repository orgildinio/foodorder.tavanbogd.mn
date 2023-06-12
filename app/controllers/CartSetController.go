package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/lambda-platform/lambda/DB"
	agentUtils "github.com/lambda-platform/lambda/agent/utils"
	"lambda/app/models"
	"net/http"
)

func AddToCartSet(c *fiber.Ctx) error {
	orderSet := models.CartMenu{}
	err := c.BodyParser(&orderSet)

	if err != nil {
		fmt.Println(err.Error())
		return c.Status(http.StatusInternalServerError).JSON("server error")
	}

	setHoolTooCartRequestData := models.SetHoolTooCartRequestData{}
	errSet := c.BodyParser(&setHoolTooCartRequestData)

	if errSet != nil {
		fmt.Println(errSet.Error())
		return c.Status(http.StatusInternalServerError).JSON("server error")
	}

	cartUser := agentUtils.AuthUserObject(c)
	orderSet.UserID = int(cartUser["id"].(int64))

	checkCart := models.CartMenuCheck{}
	DB.DB.Where("user_id = ? AND order_rule_id = ? AND age(now(), created_at) < '30 minute'", orderSet.UserID, setHoolTooCartRequestData.OrderRuleID).Order("id DESC").Find(&checkCart)

	if setHoolTooCartRequestData.Qty > 5 {
		return c.Status(http.StatusOK).JSON(map[string]interface{}{
			"status":  "warning",
			"message": "Тоо ширхэг 5-с ихгүй сонгоно уу.",
		})
	}

	if checkCart.ID > 1 {
		return c.Status(http.StatusOK).JSON(map[string]interface{}{
			"status":  "warning",
			"message": "Таньд сагсалсан хоол байна",
		})
	}

	packePrice := models.LutPacketPrice{}
	DB.DB.Find(&packePrice)

	totalPrice := int(packePrice.PacketPrice) * setHoolTooCartRequestData.Qty

	//1 save set hool
	//setHoolTooCartRequestData.MenuID

	orderSet.MenuID = setHoolTooCartRequestData.MenuID
	orderSet.OrderRuleID = setHoolTooCartRequestData.OrderRuleID
	orderSet.Qty = setHoolTooCartRequestData.Qty
	orderSet.KitchenID = setHoolTooCartRequestData.KitchenID
	orderSet.PacketPrice = totalPrice
	orderSet.IsDelivery = setHoolTooCartRequestData.IsDelivery
	orderSet.CompanyID = setHoolTooCartRequestData.CompanyID

	fmt.Println("setHoolTooCartRequestData.IsDelivery", setHoolTooCartRequestData.IsDelivery)
	fmt.Println("orderSet.IsDelivery", orderSet.IsDelivery)

	if setHoolTooCartRequestData.Qty > 5 {
		return c.Status(http.StatusOK).JSON(map[string]string{
			"status":    "warning",
			"status_mn": "Анхааруулга",
			"message":   "5-аас ихгүй сонгоно уу",
		})
	}

	DB.DB.Create(&orderSet)

	//leftTimeMenu := models.CountLeftMenuTime{}
	//
	//DB.DB.Where("id = ? AND age(morning_order_end <= 2)", orderSet.MenuID).Find(&leftTimeMenu)

	//go LeftTimeSend(leftTimeMenu.ID)

	//2 save set subs
	for _, subMenuData := range setHoolTooCartRequestData.Items {
		cartSubMenu := models.CartSubMenu{}

		cartSubMenu.MenuID = orderSet.ID
		cartSubMenu.FoodTypeID = subMenuData.FoodTypeID

		DB.DB.Create(&cartSubMenu)

		for _, subMenuFoodData := range subMenuData.SubItems {
			cartSubMenuFood := models.CartSubMenuFood{}

			cartSubMenuFood.FoodID = subMenuFoodData.FoodID
			cartSubMenuFood.SubMenuID = cartSubMenu.ID

			DB.DB.Create(&cartSubMenuFood)

		}
	}

	cartMenu := models.ViewCartMenu{}
	DB.DB.Order("id DESC").Find(&cartMenu)

	return c.Status(http.StatusOK).JSON(map[string]interface{}{
		"status":  "success",
		"message": "Сонгосон хоол сагсанд нэмэгдлээ",
		"cart_id": cartMenu.ID,
	})
}

func EditCartItem(c *fiber.Ctx) error {
	cartMenu := models.CartMenu{}
	err := c.BodyParser(&cartMenu)
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(http.StatusInternalServerError).JSON("server error")
	}

	cartUser := agentUtils.AuthUserObject(c)

	editCart := models.CartMenu{}

	DB.DB.Where("id = ? AND user_id = ?", cartMenu.ID, cartUser["id"]).Find(&editCart)

	if editCart.ID == 0 {
		return c.Status(http.StatusOK).JSON(map[string]string{
			"status":  "warning",
			"message": "Food not found",
		})
	}

	setHoolTooCartRequestData := models.SetHoolTooCartRequestData{}
	errSet := c.BodyParser(&setHoolTooCartRequestData)

	if errSet != nil {
		fmt.Println(errSet.Error())
		return c.Status(http.StatusInternalServerError).JSON("server error")
	}

	packePrice := models.LutPacketPrice{}
	DB.DB.Find(&packePrice)

	editCart.Qty = editCart.Qty + setHoolTooCartRequestData.Qty
	totalPrice := int(packePrice.PacketPrice) * editCart.Qty
	editCart.PacketPrice = totalPrice

	if 1 > editCart.Qty {
		return nil
	}

	if editCart.Qty > 5 {
		return c.Status(http.StatusOK).JSON(map[string]string{
			"status":    "warning",
			"status_mn": "Анхааруулга",
			"message":   "5-аас ихгүй сонгоно уу",
		})
	}

	DB.DB.Save(&editCart)

	return c.Status(http.StatusOK).JSON(map[string]string{
		"status":  "success",
		"message": "Update success",
	})
}

func DeleteCartItem(c *fiber.Ctx) error {
	cartMenu := models.CartMenu{}
	err := c.BodyParser(&cartMenu)

	if err != nil {
		fmt.Println(err.Error())
		return c.Status(http.StatusInternalServerError).JSON("server error")
	}

	cartUser := agentUtils.AuthUserObject(c)
	cartMenu.UserID = int(cartUser["id"].(int64))

	delCart := models.CartMenu{}
	DB.DB.Where("id = ? AND user_id = ?", cartMenu.ID, cartMenu.UserID).Find(&delCart)

	if delCart.ID == 0 {
		return c.Status(http.StatusOK).JSON(map[string]string{
			"status":  "success",
			"message": "Food not found",
		})
	}

	DB.DB.Delete(delCart)

	return c.Status(http.StatusOK).JSON(map[string]string{
		"status":  "success",
		"message": "Сагснаас хасагдлаа",
	})
}

func CheckFoodBalance(c *fiber.Ctx) error {

	setHool := []models.SetHoolTooCartRequestData{}
	fmt.Println(setHool)

	balance := []models.FoodBalance{}
	DB.DB.Where("id = ?", 25).Find(&balance)

	return c.JSON(setHool)
}
