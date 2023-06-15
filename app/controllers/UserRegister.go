package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/lambda-platform/lambda/DB"
	agentUtils "github.com/lambda-platform/lambda/agent/utils"
	"lambda/app/models"
	"net/http"
	"strings"
)

func GetIntPointer(value int) *int {
	return &value
}

func UserRegistration(c *fiber.Ctx) error {
	userRegisterData := new(models.UserRegisterRequestData)
	userData := models.Users{}

	err := c.BodyParser(&userRegisterData)

	if err != nil {
		fmt.Println(err.Error())
		return c.Status(http.StatusInternalServerError).JSON("server error")
	}

	userDataEmail := models.Users{}
	userDataPhone := models.Users{}
	userDataLogin := models.Users{}

	DB.DB.Where("email = ?", userRegisterData.Email).Find(&userDataEmail)
	DB.DB.Where("phone = ?", userRegisterData.Phone).Find(&userDataPhone)
	DB.DB.Where("login = ?", userRegisterData.Login).Find(&userDataLogin)

	if userDataEmail.ID > 0 {
		return c.Status(http.StatusOK).JSON(map[string]string{
			"status":  "error",
			"message": "Бүртгэлтэй е-майл хаяг байна.",
		})
	} else if userDataPhone.ID > 0 {
		return c.Status(http.StatusOK).JSON(map[string]string{
			"status":  "error",
			"message": "Бүртгэлтэй утасны дугаар байна.",
		})
	} else if userDataLogin.ID > 0 {
		return c.Status(http.StatusOK).JSON(map[string]string{
			"status":  "error",
			"message": "Нэвтрэх нэр бүртгэлтэй байна",
		})
	}

	checkIsSamePassword := *userRegisterData.Password == *userRegisterData.ConfirmPassword

	if checkIsSamePassword == false {
		return c.Status(http.StatusOK).JSON(map[string]string{
			"status":  "error",
			"message": "Баталгаажуулах нууц үг таарсангүй",
		})
	}

	userData.LastName = &userRegisterData.LastName
	userData.FirstName = &userRegisterData.FirstName
	userData.Phone = userRegisterData.Phone
	userData.Role = GetIntPointer(3)
	userData.Login = strings.ToLower(userRegisterData.Login)
	userData.Email = strings.ToLower(userRegisterData.Email)
	password, _ := agentUtils.Hash(*userRegisterData.Password)
	userData.Password = password

	DB.DB.Create(&userData)

	return c.Status(http.StatusOK).JSON(map[string]string{
		"status":  "success",
		"message": "Бүртгэл амжилттай",
	})
}

func UserDelete(c *fiber.Ctx) error {
	orderUser := agentUtils.AuthUserObject(c)
	user := models.DeleteUsers{}
	DB.DB.Delete(&user, "id = ?", orderUser["id"])

	return c.Status(http.StatusOK).JSON(map[string]string{
		"status":  "success",
		"message": "Хэрэглэгч устгагдлаа",
	})
}
