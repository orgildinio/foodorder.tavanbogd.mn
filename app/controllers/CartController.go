package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/lambda-platform/lambda/DB"
	agentUtils "github.com/lambda-platform/lambda/agent/utils"
	"lambda/app/models"
	"net/http"
)

func AddToCart(c *fiber.Ctx) error {
	cart := new(models.CartZahialgat)
	err := c.BodyParser(&cart)

	if err != nil {
		fmt.Println(err.Error())
		return c.Status(http.StatusInternalServerError).JSON("server error")
	}
	cartUser := agentUtils.AuthUserObject(c)

	cartFood := models.OrderFoodCarting{}
	errSet := c.BodyParser(&cartFood)
	if errSet != nil {
		fmt.Println(errSet.Error())
		return c.Status(http.StatusInternalServerError).JSON("server error")
	}

	foodBalance := models.FoodBalance{}
	DB.DB.Where("food_id = ?", cartFood.FoodID).Find(&foodBalance)

	fmt.Println("=====================>")
	fmt.Println("=====================>", foodBalance.FoodPrice)
	fmt.Println("=====================>")

	foodPrice := int(foodBalance.FoodPrice)

	totalPrice := foodPrice * cartFood.Qty

	cart.UserID = int(cartUser["id"].(int64))
	cart.FoodID = cartFood.FoodID
	cart.Price = totalPrice
	cart.Qty = cartFood.Qty

	checkCart := models.ViewCartZahialga{}
	DB.DB.Where("food_id = ? AND user_id = ? AND age(now(), created_at) < '30 minute'", checkCart.FoodID, cart.UserID).Find(&checkCart)

	if cartFood.Qty > 5 {
		return c.Status(http.StatusOK).JSON(map[string]string{
			"status":    "warning",
			"status_mn": "Анхааруулга",
			"message":   "5-с ихгүй сонгоно уу",
		})
	}

	DB.DB.Create(&cart)

	cartZahialga := models.ViewCartZahialga{}
	DB.DB.Order("id DESC").Find(&cartZahialga)

	return c.Status(http.StatusOK).JSON(map[string]interface{}{
		"status":  "success",
		"message": "Сонгосон хоол сагсан нэмэгдлээ",
		"cart_id": *cartZahialga.ID,
	})

}

func UpdateCart(c *fiber.Ctx) error {
	cartReqData := new(models.CartZahialgat)
	err := c.BodyParser(&cartReqData)
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(http.StatusInternalServerError).JSON("server error")
	}

	cartUser := agentUtils.AuthUserObject(c)

	editCart := models.CartZahialgat{}
	DB.DB.Debug().Where("id = ? AND user_id = ?", cartReqData.ID, cartUser["id"]).Find(&editCart)

	if editCart.ID == 0 {
		return c.Status(http.StatusOK).JSON(map[string]string{
			"status":  "warning",
			"message": "Food not found",
		})
	}

	editCart.Qty = editCart.Qty + cartReqData.Qty

	if cartReqData.Qty > 5 {
		return c.Status(http.StatusOK).JSON(map[string]string{
			"status":    "warning",
			"status_mn": "Анхааруулга",
			"message":   "5 аас ихгүй сонгоно уу",
		})
	}

	DB.DB.Debug().Save(&editCart)

	fmt.Println(editCart.Qty)

	return c.Status(http.StatusOK).JSON(map[string]string{
		"status":  "success",
		"message": "Update success",
	})
}

func DeleteCart(c *fiber.Ctx) error {
	cartReqData := new(models.CartZahialgat)
	err := c.BodyParser(&cartReqData)
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(http.StatusInternalServerError).JSON("server error")
	}

	cartUser := agentUtils.AuthUserObject(c)

	editCart := models.CartZahialgat{}
	DB.DB.Debug().Where("id = ? AND user_id = ?", cartReqData.ID, cartUser["id"]).Find(&editCart)

	if editCart.ID == 0 {
		return c.Status(http.StatusOK).JSON(map[string]string{
			"status":  "warning",
			"message": "Food not found",
		})
	}

	DB.DB.Delete(&editCart)

	return c.Status(http.StatusOK).JSON(map[string]string{
		"status":  "success",
		"message": "Сагсалсан хоол устгагдлаа",
	})
}
