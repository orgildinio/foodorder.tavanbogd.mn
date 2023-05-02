package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lambda-platform/lambda/DB"
	"lambda/app/models"
)

func FoodType(c *fiber.Ctx) error {

	foodData := []models.TblFood{}

	foodTypeReq := models.FoodTypeRequest{}

	if err := c.BodyParser(&foodTypeReq); err != nil {
		return err
	}

	//fmt.Println(foodTypeReq.Model.FoodTypeID)

	DB.DB.Where("food_type_id = ?", foodTypeReq.Model.FoodTypeID).Find(&foodData)

	options := []map[string]interface{}{}

	for _, food := range foodData {
		option := map[string]interface{}{
			"value": food.ID,
			"label": food.FoodName,
		}

		options = append(options, option)
	}

	SubSchemas := []interface{}{}

	field := models.Field2{
		Field: "food_id",
		Value: nil,
		Props: map[string]interface{}{
			"options": options,
		},
	}

	//field2 := models.Field{
	//	Field: "sub_menu_foods",
	//	Value: []map[string]interface{}{},
	//}

	SubSchemas = append(SubSchemas, field)

	return c.JSON(map[string]interface{}{
		"status": true,
		"schema_sub": []map[string]interface{}{
			map[string]interface{}{
				"model":  "sub_menu_foods",
				"schema": SubSchemas,
			},
		},
		"schema": []interface{}{
			//field2,
		},
	})
}
