package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/lambda-platform/lambda/DB"
	"github.com/lambda-platform/lambda/config"
	notifyHandler "github.com/lambda-platform/lambda/notify/handler"
	modelsModels "github.com/lambda-platform/lambda/notify/models"
	"lambda/app/models"
)

func FoodBalance(c *fiber.Ctx) error {
	var balance []models.FoodBalance
	DB.DB.Find(&balance)

	//LeftOverQuantity()

	return c.JSON(balance)
}

func LeftOverQuantity(c *fiber.Ctx) error {

	FCMData := modelsModels.FCMData{
		Title:       "Мэдэгдэл",
		Body:        "Мэдэгдлийн жишээ, админ нүүр хуудас харна уу",
		FirstName:   "Мөнх-Алтай",
		Sound:       config.LambdaConfig.Notify.Sound,
		Icon:        config.LambdaConfig.Favicon,
		Link:        "/admin",
		ClickAction: config.LambdaConfig.Domain + "/admin",
	}

	FCMNotification := modelsModels.FCMNotification{
		Title:       "Мэдэгдэл",
		Body:        "Мэдэгдлийн жишээ, админ нүүр хуудас харна уу",
		Icon:        config.LambdaConfig.Domain + "/" + config.LambdaConfig.Favicon,
		ClickAction: config.LambdaConfig.Domain + "/admin",
	}
	data := modelsModels.NotificationData{
		Users: []int{17},
		//Roles:        []int{2},
		Data:         FCMData,
		Notification: FCMNotification,
	}
	notifyHandler.CreateNotification(data, map[string]interface{}{})
	return c.JSON(data)
}

func LeftOverQuantitySend(FoodID int, KitchenID int, Quantity int) {

	food := models.TblFood{}

	DB.DB.Where("food_id = ? AND kitchen_id = ?", FoodID, KitchenID).Find(&food)

	FCMData := modelsModels.FCMData{
		Title:       "Хоолны үлдэгдэл дуусж байна",
		Body:        fmt.Sprintf("%s гал тогооны, %s хоолны үлдэгдэл дуусж байна. Үлдэгдэл %d !!!", "1-р", food.FoodName, Quantity),
		FirstName:   "Систем",
		Sound:       config.LambdaConfig.Notify.Sound,
		Icon:        config.LambdaConfig.Favicon,
		Link:        "/admin",
		ClickAction: config.LambdaConfig.Domain + "/admin",
	}

	FCMNotification := modelsModels.FCMNotification{
		Title:       "Хоолны үлдэгдэл дуусж байна",
		Body:        fmt.Sprintf("%s гал тогооны, %s хоолны үлдэгдэл дуусж байна. Үлдэгдэл %d !!!", "1-р", food.FoodName, Quantity),
		Icon:        config.LambdaConfig.Domain + "/" + config.LambdaConfig.Favicon,
		ClickAction: config.LambdaConfig.Domain + "/admin",
	}

	data := modelsModels.NotificationData{
		//Users: []int{17},
		Roles:        []int{3},
		Data:         FCMData,
		Notification: FCMNotification,
	}
	notifyHandler.CreateNotification(data, map[string]interface{}{})

}
