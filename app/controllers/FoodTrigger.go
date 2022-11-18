package controllers

import (
	"github.com/lambda-platform/lambda/DB"
	"lambda/app/models"
	"lambda/lambda/models/form/formModels"
)

func MenuAfterInsertUpdate(menuPre interface{}) {
	var menu *formModels.TblMenu21 = menuPre.(*formModels.TblMenu21)

	uuhYums := []models.SubMenuUuhim{}
	deserts := []models.SubMenuDesert{}
	negHools := []models.SubMenuNeg{}
	hoerHools := []models.SubMenuHoer{}
	salats := []models.SubMenuSalat{}

	DB.DB.Where("menu_form_id = ?", menu.ID).Find(&uuhYums)
	DB.DB.Where("menu_form_id = ?", menu.ID).Find(&deserts)
	DB.DB.Where("menu_form_id = ?", menu.ID).Find(&negHools)
	DB.DB.Where("menu_form_id = ?", menu.ID).Find(&hoerHools)
	DB.DB.Where("menu_form_id = ?", menu.ID).Find(&salats)

	SaveGalTogoo1(menu, uuhYums, deserts, negHools, hoerHools, salats)
	//SaveGalTogoo2(menu, uuhYums, deserts)
	//SaveGalTogoo3(menu, uuhYums, deserts)

}

func SaveGalTogoo1(menu *formModels.TblMenu21, uuhYums []models.SubMenuUuhim, deserts []models.SubMenuDesert, negHools []models.SubMenuNeg, hoerHools []models.SubMenuHoer, salats []models.SubMenuSalat) {

	galTogoo1 := models.TblMenuGtNeg{}
	DB.DB.Where("main_menu_id = ?", menu.ID).Find(&galTogoo1)
	if galTogoo1.ID >= 1 {
		//DO update

		existingUuhYUmsIDS := []int{}
		for _, uuhYum := range uuhYums {
			gtUuhYum := models.SubMenuUuhimGtNeg{}

			DB.DB.Where("food_uuhim_id = ? AND menu_form_id = ?", uuhYum.FoodUuhimID, galTogoo1.ID).Find(&gtUuhYum)

			if gtUuhYum.ID <= 0 {
				gtUuhYum.MenuFormID = &galTogoo1.ID
				gtUuhYum.FoodUuhimID = uuhYum.FoodUuhimID
				DB.DB.Create(&gtUuhYum)
			}

			existingUuhYUmsIDS = append(existingUuhYUmsIDS, gtUuhYum.ID)

		}

		DB.DB.Where("menu_form_id = ? AND id NOT IN ?", galTogoo1.ID, existingUuhYUmsIDS).Delete(models.SubMenuUuhimGtNeg{})

		existingDesertIDS := []int{}
		for _, desert := range deserts {
			gtDesert := models.SubMenuDesertGtNeg{}

			DB.DB.Where("food_desert_id = ? AND menu_form_id = ?", desert.FoodDesertID, galTogoo1.ID).Find(&gtDesert)

			if gtDesert.ID <= 0 {
				gtDesert.MenuFormID = &galTogoo1.ID
				gtDesert.FoodDesertID = desert.FoodDesertID
				DB.DB.Create(&gtDesert)
			}
			existingDesertIDS = append(existingDesertIDS, gtDesert.ID)
		}

		DB.DB.Where("menu_form_id = ? AND id NOT IN ?", galTogoo1.ID, existingDesertIDS).Delete(models.SubMenuDesertGtNeg{})

		existingNegHoolsIDS := []int{}
		for _, negHool := range negHools {
			gtNegHool := models.SubMenuNegGtNeg{}

			DB.DB.Where("food_neg_id = ? AND menu_form_id = ?", negHool.FoodNegID, galTogoo1.ID).Find(&gtNegHool)

			if gtNegHool.ID <= 0 {
				gtNegHool.MenuFormID = &galTogoo1.ID
				gtNegHool.FoodNegID = negHool.FoodNegID
				DB.DB.Create(&gtNegHool)
			}

			existingNegHoolsIDS = append(existingNegHoolsIDS, gtNegHool.ID)

		}

		DB.DB.Where("menu_form_id = ? AND id NOT IN ?", galTogoo1.ID, existingNegHoolsIDS).Delete(models.SubMenuNegGtNeg{})

		existingHoerHoolIDS := []int{}
		for _, hoerHool := range hoerHools {
			gtHoerHool := models.SubMenuHoerGtNeg{}

			DB.DB.Where("food_hoer_id = ? AND menu_form_id = ?", hoerHool.FoodHoerID, galTogoo1.ID).Find(&gtHoerHool)

			if gtHoerHool.ID <= 0 {
				gtHoerHool.MenuFormID = &galTogoo1.ID
				gtHoerHool.FoodHoerID = hoerHool.FoodHoerID
				DB.DB.Create(&gtHoerHool)
			}

			existingHoerHoolIDS = append(existingHoerHoolIDS, gtHoerHool.ID)
		}

		DB.DB.Where("menu_form_id = ? AND id NOT IN ?", galTogoo1.ID, existingHoerHoolIDS).Delete(models.SubMenuHoerGtNeg{})

		existingSalatIDS := []int{}
		for _, salat := range salats {
			gtSalat := models.SubMenuSalatGtNeg{}

			DB.DB.Where("food_salat_id = ? AND menu_form_id = ?", salat.FoodSalatID, galTogoo1.ID).Find(&salat)

			if gtSalat.ID <= 0 {
				gtSalat.MenuFormID = &galTogoo1.ID
				gtSalat.FoodSalatID = salat.FoodSalatID
				DB.DB.Create(&gtSalat)
			}

			existingSalatIDS = append(existingSalatIDS, gtSalat.ID)
		}

		DB.DB.Where("menu_form_id = ? AND id NOT IN ?", galTogoo1.ID, existingSalatIDS).Delete(models.SubMenuSalatGtNeg{})

	} else {
		galTogoo1.MainMenuID = &menu.ID
		galTogoo1.SetDate = menu.SetDate
		galTogoo1.SetName = menu.SetName
		galTogoo1.FoodTimeTypeID = menu.FoodTimeTypeID

		DB.DB.Create(&galTogoo1)

		for _, uuhYum := range uuhYums {
			gtUuhYum := models.SubMenuUuhimGtNeg{}

			gtUuhYum.MenuFormID = &galTogoo1.ID
			gtUuhYum.FoodUuhimID = uuhYum.FoodUuhimID

			DB.DB.Create(&gtUuhYum)
		}

		for _, desert := range deserts {
			gtDesert := models.SubMenuDesertGtNeg{}

			gtDesert.MenuFormID = &galTogoo1.ID
			gtDesert.FoodDesertID = desert.FoodDesertID

			DB.DB.Create(&gtDesert)
		}

		for _, negHool := range negHools {
			gtNegHool := models.SubMenuNegGtNeg{}

			gtNegHool.MenuFormID = &galTogoo1.ID
			gtNegHool.FoodNegID = negHool.FoodNegID

			DB.DB.Create(&gtNegHool)
		}

		for _, hoerHool := range hoerHools {
			gtHoerHool := models.SubMenuHoerGtNeg{}

			gtHoerHool.MenuFormID = &galTogoo1.ID
			gtHoerHool.FoodHoerID = hoerHool.FoodHoerID

			DB.DB.Create(&gtHoerHool)
		}

	}
}
