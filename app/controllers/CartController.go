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
	cart.UserID = int(cartUser["id"].(int64))

	DB.DB.Create(&cart)
	return c.Status(http.StatusOK).JSON(map[string]interface{}{
		"status":  "success",
		"message": "Сагсанд нэмэгдлээ",
	})

}

func UpdateCart(c *fiber.Ctx) error {
	cartReqData := new(models.CartZahialgat)
	err := c.BodyParser(&cartReqData)
	if err != nil {
		fmt.Errorf(err.Error())
		return c.Status(http.StatusInternalServerError).JSON("server error")
	}

	cartUser := agentUtils.AuthUserObject(c)

	editCart := models.CartZahialgat{}
	DB.DB.Debug().Where("id = ? AND user_id = ?", cartReqData.ID, cartUser["id"]).Find(&editCart)

	//if editCart.ID == 0 {
	//	return c.Status(http.StatusOK).JSON(map[string]string{
	//		"status":  "warning",
	//		"message": "Food not found",
	//	})
	//}

	editCart.Qty = cartReqData.Qty

	fmt.Println(editCart.Qty)

	if cartReqData.Qty >= 6 {
		return c.Status(http.StatusOK).JSON(map[string]string{
			"status":    "warning",
			"status_mn": "Анхааруулга",
			"message":   "5 хоол захиалах боломжтой",
		})
	}

	DB.DB.Debug().Save(&editCart)

	return c.Status(http.StatusOK).JSON(map[string]string{
		"status":  "success",
		"message": "Update success",
	})
}

func DeleteCart(c *fiber.Ctx) error {
	cartReqData := new(models.CartZahialgat)
	err := c.BodyParser(&cartReqData)
	if err != nil {
		fmt.Errorf(err.Error())
		return c.Status(http.StatusInternalServerError).JSON("server error")
	}

	cartUser := agentUtils.AuthUserObject(c)

	editCart := models.CartZahialgat{}
	DB.DB.Debug().Where("id = ? AND user_id = ?", cartReqData.ID, cartUser["id"]).Find(&editCart)

	//if editCart.ID == 0 {
	//	return c.Status(http.StatusOK).JSON(map[string]string{
	//		"status":  "warning",
	//		"message": "Food not found",
	//	})
	//}

	DB.DB.Delete(&editCart)

	return c.Status(http.StatusOK).JSON(map[string]string{
		"status":  "success",
		"message": "Сагсалсан хоол устгагдлаа",
	})
}
