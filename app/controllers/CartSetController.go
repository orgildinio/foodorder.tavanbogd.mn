package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/lambda-platform/lambda/DB"
	agentUtils "github.com/lambda-platform/lambda/agent/utils"
	"lambda/app/models"
	"net/http"
)

func AddCartSet(c *fiber.Ctx) error {
	setCart := new(models.CartSet)
	err := c.BodyParser(&setCart)

	if err != nil {
		fmt.Println(err.Error())
		return c.Status(http.StatusInternalServerError).JSON("server error")
	}

	checkCart := models.ViewCartSet{}
	DB.DB.Where("user_id = ? AND menu_id = ? AND sub_menu_id = ? AND sub_menu_food_id = ? AND food_id = ?", setCart.UserID, setCart.MenuID, setCart.SubMenuID, setCart.SubMenuFoodID, setCart.FoodID).Find(&checkCart)

	if setCart.Qty > 5 {
		return c.Status(http.StatusOK).JSON(map[string]interface{}{
			"status":  "warning",
			"message": "захиалга 5-аас ихгүй байна.",
		})
	}

	if checkCart.ID >= 1 {
		return c.Status(http.StatusOK).JSON(map[string]interface{}{
			"status":  "warning",
			"message": "Таньд сагсалсан хоол байна.",
			"data":    checkCart.SetName,
		})
	} else {
		DB.DB.Create(&setCart)

		return c.Status(http.StatusOK).JSON(map[string]interface{}{
			"status":  "success",
			"message": "Сагсанд нэмэгдлээ",
			"data":    checkCart.SetName,
		})
	}
}

func UpdateCartSet(c *fiber.Ctx) error {
	cartReqData := new(models.ViewCartSet)
	err := c.BodyParser(&cartReqData)

	if err != nil {
		fmt.Errorf(err.Error())
		return c.Status(http.StatusInternalServerError).JSON("server error")
	}

	editCart := models.CartSet{}
	DB.DB.Debug().Where("id = ? AND user_id = ?", cartReqData.ID, cartReqData.UserID).Find(&editCart)

	editCart.Qty = editCart.Qty + cartReqData.Qty

	fmt.Println(editCart.Qty)

	if editCart.Qty > 5 {
		return c.Status(http.StatusOK).JSON(map[string]string{
			"status":    "warning",
			"status_mn": "Анхааруулга",
			"message":   "Сет хоол 5-г захиалах боломжтой",
		})
	}

	DB.DB.Debug().Save(&editCart)

	return c.Status(http.StatusOK).JSON(map[string]string{
		"status":  "success",
		"message": "Сагсалсан хоол засагдлаа",
	})
}

func DeleteCartSet(c *fiber.Ctx) error {
	cartReqData := new(models.ViewCartSet)
	err := c.BodyParser(&cartReqData)
	if err != nil {
		fmt.Errorf(err.Error())
		return c.Status(http.StatusInternalServerError).JSON("server error")
	}

	cartUser := agentUtils.AuthUserObject(c)

	editCart := models.CartSet{}
	DB.DB.Debug().Where("id = ? AND user_id = ?", cartReqData.ID, cartUser["id"]).Delete(&editCart)

	//DB.DB.Debug().Delete(&editCart)

	return c.Status(http.StatusOK).JSON(map[string]string{
		"status":  "success",
		"message": "Мэдээлэл устгагдлаа",
	})
}
