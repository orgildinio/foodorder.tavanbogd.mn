package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/lambda-platform/lambda/DB"
	"gorm.io/gorm"
	"lambda/app/models"
	"net/http"
	"time"
)

func CreateOrder(c *fiber.Ctx) error {

	var activeMenu []models.AcitiveMenu

	emptyCheck := DB.DB.First(&activeMenu).Error

	if emptyCheck == gorm.ErrRecordNotFound {
		return c.Status(http.StatusNotFound).JSON(map[string]string{
			"status":  "unsuccessful",
			"message": "Идэвхтэй цэс байхгүй байна",
		})

	} else {
		orderReq := new(models.Orders)
		orderErr := c.BodyParser(&orderReq)

		if orderErr != nil {
			fmt.Println(orderErr.Error())
			return c.Status(http.StatusInternalServerError).JSON("server error")
		}

		var activeOrder []models.ActiveOrder

		DB.DB.Table("active_order").Where("user_id = ?", orderReq.UserID).Find(&activeOrder)

		if len(activeOrder) > 0 {
			return c.Status(http.StatusOK).JSON(map[string]string{
				"status":  "warning",
				"message": "Таньд захиалга байна",
			})
		}

		DB.DB.Create(&orderReq)
		return c.Status(http.StatusOK).JSON(map[string]string{
			"status":  "success",
			"message": "Захиалга амжилттай үүсгэлээ",
		})
	}

}

func TimeCounter(c *fiber.Ctx) error {

	activeOrder := models.ActiveOrder{}
	DB.DB.Where("order_status = ?", "pending").Find(&activeOrder)

	curActive := activeOrder.CancelledAt.Format("2006-01-02 15:04:05")

	fmt.Println("Hey", curActive)

	for {
		curTime := time.Now().Format("2006-01-02 15:04:05")
		time.Sleep(1 * time.Second)
		fmt.Println(curTime)

		if curActive == curTime {
			orders := models.Orders{}
			DB.DB.Where("order_status = ?", "pending").Find(&orders)
			orders.OrderStatus = "update with TIMESTAMP"
			DB.DB.Save(&orders)
			fmt.Println("Hello World")
		}
	}

	return c.JSON("Hello")
}
