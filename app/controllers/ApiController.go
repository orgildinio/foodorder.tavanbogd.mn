package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/lambda-platform/lambda/DB"
	"github.com/lambda-platform/lambda/config"
	notifyHandler "github.com/lambda-platform/lambda/notify/handler"
	modelsModels "github.com/lambda-platform/lambda/notify/models"
	"lambda/app/models"
	"os"
	"time"
)

func Users(c *fiber.Ctx) error {

	users := []models.Users{}
	DB.DB.Find(&users)

	return c.JSON(users)
}

func SendOrderNotification(c *fiber.Ctx) error {

	TimeTick()
	now := time.Now()
	formattedTimeStamp := now.Format("2006-01-02 15:04:05")
	fmt.Println(formattedTimeStamp)
	fmt.Println("Sent Order Notification")

	return c.JSON(map[string]interface{}{
		"msg": "sent",
	})
}

func SendNotification(c *fiber.Ctx) error {

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

func ReadIcons(c *fiber.Ctx) error {

	files, err := os.ReadDir("public/assets/icons/duotune/")
	if err != nil {
		fmt.Println(err)
	}

	svgIcons := []map[string]interface{}{}
	for _, file := range files {
		if file.IsDir() {
			svgGroup := map[string]interface{}{}

			svgGroup["title"] = "SVG Duotune " + file.Name()

			iconList := []map[string]string{}
			svgs, err2 := os.ReadDir("public/assets/icons/duotune/" + file.Name())
			if err2 != nil {
				fmt.Println(err2)
			} else {
				for _, svg := range svgs {
					iconList = append(iconList, map[string]string{
						"url":  "/assets/icons/duotune/" + file.Name() + "/" + svg.Name(),
						"name": svg.Name(),
					})
				}
			}
			svgGroup["svgs"] = iconList

			//fmt.Println(file.Name(), file.IsDir())

			svgIcons = append(svgIcons, svgGroup)
		}

	}
	return c.JSON(svgIcons)
}
