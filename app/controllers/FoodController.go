package controllers

import (
	"github.com/lambda-platform/lambda/DB"
	"lambda/app/models"
	"lambda/lambda/models/form/formModels"
)

func InsertFoodBalance(foodPre interface{}) {
	var food *formModels.TblFood18 = foodPre.(*formModels.TblFood18)
	DB.DB.Find(&food)

	var kitchens []models.Kitchen
	DB.DB.Find(&kitchens)

	for _, kitchen := range kitchens {

		foodBalance := models.FoodBalance{}
		DB.DB.Debug().Where("food_id = ? AND kitchen_id = ?", food.ID, kitchen.ID).Find(&foodBalance)

		if foodBalance.ID == 0 {
			foodBalance.KitchenID = kitchen.ID
			foodBalance.FoodID = food.ID

			DB.DB.Debug().Create(&foodBalance)
		}

	}
}
