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
	user := agentUtils.AuthUserObject(c)
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

	foodBalance := models.ViewFoodBalance{}
	DB.DB.Where("food_id = ? AND kitchen_id = ?", cartFood.FoodID, cart.KitchenID).Find(&foodBalance)

	if foodBalance.Quantity < cartFood.Qty {
		return c.Status(http.StatusOK).JSON(map[string]string{
			"status":    "warning",
			"status_mn": "Анхааруулга",
			"message":   "Сонгосон хоолны үлдэгдэл хүрэлцэхгүй байна",
		})
	}

	if 1 > cartFood.Qty {
		return c.Status(http.StatusOK).JSON(map[string]string{
			"status":    "warning",
			"status_mn": "Анхааруулга",
			"message":   "1-с багагүй сонгоно уу",
		})
	}

	foodPrice := int(foodBalance.FoodPrice)
	totalPrice := foodPrice * cartFood.Qty

	cart.UserID = int(cartUser["id"].(int64))
	cart.FoodID = cartFood.FoodID
	cart.Price = totalPrice
	cart.Qty = cartFood.Qty
	cart.IsDelivery = cartFood.IsDelivery
	cart.CompanyID = cartFood.CompanyID

	checkCart := models.ViewCartZahialga{}
	DB.DB.Where("food_id = ? AND user_id = ?", checkCart.FoodID, user["id"]).Find(&checkCart)

	if cartFood.Qty > 5 {
		return c.Status(http.StatusOK).JSON(map[string]string{
			"status":    "warning",
			"status_mn": "Анхааруулга",
			"message":   "5-с ихгүй сонгоно уу",
		})
	}

	DB.DB.Debug().Create(&cart)

	cartZahialga := models.ViewCartZahialga{}
	DB.DB.Order("id DESC").Find(&cartZahialga)

	return c.Status(http.StatusOK).JSON(map[string]interface{}{
		"status":  "success",
		"message": "Сонгосон хоол сагсан нэмэгдлээ",
		"cart_id": cartZahialga.ID,
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

	foodBalance := models.ViewFoodBalance{}
	DB.DB.Where("food_id = ? AND kitchen_id = ?", editCart.FoodID, editCart.KitchenID).Find(&foodBalance)

	editCart.Qty = editCart.Qty + cartReqData.Qty
	editCart.Price = int(foodBalance.FoodPrice) * editCart.Qty

	if foodBalance.Quantity < editCart.Qty {
		return c.Status(http.StatusOK).JSON(map[string]string{
			"status":    "warning",
			"status_mn": "Анхааруулга",
			"message":   "Сонгосон хоолны үлдэгдэл хүрэлцэхгүй байна",
		})
	}

	if 1 > editCart.Qty {
		return c.Status(http.StatusOK).JSON(map[string]string{
			"status":    "warning",
			"status_mn": "Анхааруулга",
			"message":   "1 - с багагүй байна",
		})
	}

	if 5 < editCart.Qty {
		return c.Status(http.StatusOK).JSON(map[string]string{
			"status":    "warning",
			"status_mn": "Анхааруулга",
			"message":   "5 - аас ихгүй сонгоно уу",
		})
	}

	DB.DB.Save(&editCart)

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
