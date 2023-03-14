package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/lambda-platform/lambda/DB"
	"lambda/app/models"
	"net/http"
)

func CreateOrderSet(c *fiber.Ctx) error {
	orderSet := models.OrderSet{}
	err := c.BodyParser(&orderSet)

	if err != nil {
		fmt.Println(err.Error())
		return c.Status(http.StatusInternalServerError).JSON("server error")
	}

	setMenu := models.ViewCartSetFood{}
	DB.DB.Where("menu_id = ? AND id = ?", orderSet.MenuID, orderSet.CartID).Find(&setMenu)

	if setMenu.ID < 1 {
		return c.Status(http.StatusOK).JSON(map[string]interface{}{
			"status":  "success",
			"message": "Сагсалсан хоол олдсонгүй",
		})
	}

	orderSet.MenuID = setMenu.MenuID
	orderSet.CartID = setMenu.ID
	orderSet.UserID = setMenu.UserID
	orderSet.OrderStatus = GetStringPointer("pending")

	subMenus := []models.ViewSetFoodSubMenu{}
	for _, subMenu := range subMenus {
		cartSubMenus := models.CartSubMenu{}
		//DB.DB.Find(&cartSubMenus)

		cartSubMenus.MenuID = subMenu.MenuID
		cartSubMenus.FoodTypeID = subMenu.FoodTypeID

		DB.DB.Create(&cartSubMenus)

		subMenuFoods := []models.ViewSetFoodSubMenuFoods{}
		for _, subMenuFood := range subMenuFoods {
			cartSubMenuFood := models.CartSubMenuFood{}
			//DB.DB.Find(&cartSubMenuFood)

			cartSubMenuFood.SubMenuID = subMenuFood.SubMenuID
			cartSubMenuFood.FoodID = subMenuFood.FoodID

			DB.DB.Create(cartSubMenuFood)

		}
	}

	orderCheck := models.OrderSetCheck{}
	DB.DB.Where("user_id = ? AND menu_id = ? AND cart_id = ? AND age(created_at) <= '15 minutes' AND order_status = 'pending'", orderSet.UserID, orderSet.MenuID, orderSet.CartID).Order("id DESC").Find(&orderCheck)

	if orderCheck.ID >= 1 {
		return c.JSON(map[string]string{
			"status":      "warning",
			"orderNumber": "Таньд " + *orderCheck.OrderNumber + " дугаартай захиалга үүссэн байна.",
			"status_mn":   "Анхааруулга",
		})
	}

	DB.DB.Create(&orderSet)

	return c.Status(http.StatusOK).JSON(map[string]interface{}{
		"status":  "success",
		"message": "Захиалга үүсгэлээ.",
	})
}
