package controllers

import (
	"fmt"
	"lambda/app/models"
	"lambda/lambda/models/form/formModels"

	"github.com/lambda-platform/lambda/DB"
)

//import (
//	"github.com/lambda-platform/lambda/DB"
//	"lambda/app/models"
//	"lambda/lambda/models/form/formModels"
//)
//
//func MenuAfterInsertUpdate(menuPre interface{}) {
//	var menu *formModels.TblMenu21 = menuPre.(*formModels.TblMenu21)
//
//	uuhYums := []models.SubMenuUuhim{}
//	deserts := []models.SubMenuDesert{}
//	negHools := []models.SubMenuNeg{}
//	hoerHools := []models.SubMenuHoer{}
//	salats := []models.SubMenuSalat{}
//
//	DB.DB.Where("menu_form_id = ?", menu.ID).Find(&uuhYums)
//	DB.DB.Where("menu_form_id = ?", menu.ID).Find(&deserts)
//	DB.DB.Where("menu_form_id = ?", menu.ID).Find(&negHools)
//	DB.DB.Where("menu_form_id = ?", menu.ID).Find(&hoerHools)
//	DB.DB.Where("menu_form_id = ?", menu.ID).Find(&salats)
//
//	SaveGalTogoo1(menu, uuhYums, deserts, negHools, hoerHools, salats)
//	SaveGalTogoo2(menu, uuhYums, deserts, negHools, hoerHools, salats)
//	SaveGalTogoo3(menu, uuhYums, deserts, negHools, hoerHools, salats)
//
//}
//
//func SaveGalTogoo1(menu *formModels.TblMenu21, uuhYums []models.SubMenuUuhim, deserts []models.SubMenuDesert, negHools []models.SubMenuNeg, hoerHools []models.SubMenuHoer, salats []models.SubMenuSalat) {
//
//	galTogoo1 := models.TblMenuGtNeg{}
//	DB.DB.Where("main_menu_id = ?", menu.ID).Find(&galTogoo1)
//	if galTogoo1.ID >= 1 {
//		//DO update
//
//		existingUuhYUmsIDS := []int{}
//		for _, uuhYum := range uuhYums {
//			gtUuhYum := models.SubMenuUuhimGtNeg{}
//
//			DB.DB.Where("food_uuhim_id = ? AND menu_form_id = ?", uuhYum.FoodUuhimID, galTogoo1.ID).Find(&gtUuhYum)
//
//			if gtUuhYum.ID <= 0 {
//				gtUuhYum.MenuFormID = &galTogoo1.ID
//				gtUuhYum.FoodUuhimID = uuhYum.FoodUuhimID
//				DB.DB.Create(&gtUuhYum)
//			}
//
//			existingUuhYUmsIDS = append(existingUuhYUmsIDS, gtUuhYum.ID)
//
//		}
//
//		DB.DB.Where("menu_form_id = ? AND id NOT IN ?", galTogoo1.ID, existingUuhYUmsIDS).Delete(models.SubMenuUuhimGtNeg{})
//
//		existingDesertIDS := []int{}
//		for _, desert := range deserts {
//			gtDesert := models.SubMenuDesertGtNeg{}
//
//			DB.DB.Where("food_desert_id = ? AND menu_form_id = ?", desert.FoodDesertID, galTogoo1.ID).Find(&gtDesert)
//
//			if gtDesert.ID <= 0 {
//				gtDesert.MenuFormID = &galTogoo1.ID
//				gtDesert.FoodDesertID = desert.FoodDesertID
//				DB.DB.Create(&gtDesert)
//			}
//			existingDesertIDS = append(existingDesertIDS, gtDesert.ID)
//		}
//
//		DB.DB.Where("menu_form_id = ? AND id NOT IN ?", galTogoo1.ID, existingDesertIDS).Delete(models.SubMenuDesertGtNeg{})
//
//		existingNegHoolsIDS := []int{}
//		for _, negHool := range negHools {
//			gtNegHool := models.SubMenuNegGtNeg{}
//
//			DB.DB.Where("food_neg_id = ? AND menu_form_id = ?", negHool.FoodNegID, galTogoo1.ID).Find(&gtNegHool)
//
//			if gtNegHool.ID <= 0 {
//				gtNegHool.MenuFormID = &galTogoo1.ID
//				gtNegHool.FoodNegID = negHool.FoodNegID
//				DB.DB.Create(&gtNegHool)
//			}
//
//			existingNegHoolsIDS = append(existingNegHoolsIDS, gtNegHool.ID)
//
//		}
//
//		DB.DB.Where("menu_form_id = ? AND id NOT IN ?", galTogoo1.ID, existingNegHoolsIDS).Delete(models.SubMenuNegGtNeg{})
//
//		existingHoerHoolIDS := []int{}
//		for _, hoerHool := range hoerHools {
//			gtHoerHool := models.SubMenuHoerGtNeg{}
//
//			DB.DB.Where("food_hoer_id = ? AND menu_form_id = ?", hoerHool.FoodHoerID, galTogoo1.ID).Find(&gtHoerHool)
//
//			if gtHoerHool.ID <= 0 {
//				gtHoerHool.MenuFormID = &galTogoo1.ID
//				gtHoerHool.FoodHoerID = hoerHool.FoodHoerID
//				DB.DB.Create(&gtHoerHool)
//			}
//
//			existingHoerHoolIDS = append(existingHoerHoolIDS, gtHoerHool.ID)
//		}
//
//		DB.DB.Where("menu_form_id = ? AND id NOT IN ?", galTogoo1.ID, existingHoerHoolIDS).Delete(models.SubMenuHoerGtNeg{})
//
//		existingSalatIDS := []int{}
//		for _, salat := range salats {
//			gtSalat := models.SubMenuSalatGtNeg{}
//
//			DB.DB.Where("food_salat_id = ? AND menu_form_id = ?", salat.FoodSalatID, galTogoo1.ID).Find(&salat)
//
//			if gtSalat.ID <= 0 {
//				gtSalat.MenuFormID = &galTogoo1.ID
//				gtSalat.FoodSalatID = salat.FoodSalatID
//				DB.DB.Create(&gtSalat)
//			}
//
//			existingSalatIDS = append(existingSalatIDS, gtSalat.ID)
//		}
//
//		DB.DB.Where("menu_form_id = ? AND id NOT IN ?", galTogoo1.ID, existingSalatIDS).Delete(models.SubMenuSalatGtNeg{})
//
//	} else {
//		galTogoo1.MainMenuID = &menu.ID
//		galTogoo1.SetDate = menu.SetDate
//		galTogoo1.SetName = menu.SetName
//		galTogoo1.FoodTimeTypeID = menu.FoodTimeTypeID
//
//		DB.DB.Create(&galTogoo1)
//
//		for _, uuhYum := range uuhYums {
//			gtUuhYum := models.SubMenuUuhimGtNeg{}
//
//			gtUuhYum.MenuFormID = &galTogoo1.ID
//			gtUuhYum.FoodUuhimID = uuhYum.FoodUuhimID
//
//			DB.DB.Create(&gtUuhYum)
//		}
//
//		for _, desert := range deserts {
//			gtDesert := models.SubMenuDesertGtNeg{}
//
//			gtDesert.MenuFormID = &galTogoo1.ID
//			gtDesert.FoodDesertID = desert.FoodDesertID
//
//			DB.DB.Create(&gtDesert)
//		}
//
//		for _, negHool := range negHools {
//			gtNegHool := models.SubMenuNegGtNeg{}
//
//			gtNegHool.MenuFormID = &galTogoo1.ID
//			gtNegHool.FoodNegID = negHool.FoodNegID
//
//			DB.DB.Create(&gtNegHool)
//		}
//
//		for _, hoerHool := range hoerHools {
//			gtHoerHool := models.SubMenuHoerGtNeg{}
//
//			gtHoerHool.MenuFormID = &galTogoo1.ID
//			gtHoerHool.FoodHoerID = hoerHool.FoodHoerID
//
//			DB.DB.Create(&gtHoerHool)
//		}
//
//	}
//}
//
//func SaveGalTogoo2(menu *formModels.TblMenu21, uuhYums []models.SubMenuUuhim, deserts []models.SubMenuDesert, negHools []models.SubMenuNeg, hoerHools []models.SubMenuHoer, salats []models.SubMenuSalat) {
//
//	galTogoo2 := models.TblMenuGtHoer{}
//	DB.DB.Where("main_menu_id = ?", menu.ID).Find(&galTogoo2)
//	if galTogoo2.ID >= 1 {
//		//DO update
//
//		existingUuhYUmsIDS := []int{}
//		for _, uuhYum := range uuhYums {
//			gtUuhYum := models.SubMenuUuhimGtHoer{}
//
//			DB.DB.Where("food_uuhim_id = ? AND menu_form_id = ?", uuhYum.FoodUuhimID, galTogoo2.ID).Find(&gtUuhYum)
//
//			if gtUuhYum.ID <= 0 {
//				gtUuhYum.MenuFormID = &galTogoo2.ID
//				gtUuhYum.FoodUuhimID = uuhYum.FoodUuhimID
//				DB.DB.Create(&gtUuhYum)
//			}
//
//			existingUuhYUmsIDS = append(existingUuhYUmsIDS, gtUuhYum.ID)
//
//		}
//
//		DB.DB.Where("menu_form_id = ? AND id NOT IN ?", galTogoo2.ID, existingUuhYUmsIDS).Delete(models.SubMenuUuhimGtHoer{})
//
//		existingDesertIDS := []int{}
//		for _, desert := range deserts {
//			gtDesert := models.SubMenuDesertGtHoer{}
//
//			DB.DB.Where("food_desert_id = ? AND menu_form_id = ?", desert.FoodDesertID, galTogoo2.ID).Find(&gtDesert)
//
//			if gtDesert.ID <= 0 {
//				gtDesert.MenuFormID = &galTogoo2.ID
//				gtDesert.FoodDesertID = desert.FoodDesertID
//				DB.DB.Create(&gtDesert)
//			}
//			existingDesertIDS = append(existingDesertIDS, gtDesert.ID)
//		}
//
//		DB.DB.Where("menu_form_id = ? AND id NOT IN ?", galTogoo2.ID, existingDesertIDS).Delete(models.SubMenuDesertGtHoer{})
//
//		existingNegHoolsIDS := []int{}
//		for _, negHool := range negHools {
//			gtNegHool := models.SubMenuNegGtHoer{}
//
//			DB.DB.Where("food_neg_id = ? AND menu_form_id = ?", negHool.FoodNegID, galTogoo2.ID).Find(&gtNegHool)
//
//			if gtNegHool.ID <= 0 {
//				gtNegHool.MenuFormID = &galTogoo2.ID
//				gtNegHool.FoodNegID = negHool.FoodNegID
//				DB.DB.Create(&gtNegHool)
//			}
//
//			existingNegHoolsIDS = append(existingNegHoolsIDS, gtNegHool.ID)
//
//		}
//
//		DB.DB.Where("menu_form_id = ? AND id NOT IN ?", galTogoo2.ID, existingNegHoolsIDS).Delete(models.SubMenuNegGtHoer{})
//
//		existingHoerHoolIDS := []int{}
//		for _, hoerHool := range hoerHools {
//			gtHoerHool := models.SubMenuHoerGtHoer{}
//
//			DB.DB.Where("food_hoer_id = ? AND menu_form_id = ?", hoerHool.FoodHoerID, galTogoo2.ID).Find(&gtHoerHool)
//
//			if gtHoerHool.ID <= 0 {
//				gtHoerHool.MenuFormID = &galTogoo2.ID
//				gtHoerHool.FoodHoerID = hoerHool.FoodHoerID
//				DB.DB.Create(&gtHoerHool)
//			}
//
//			existingHoerHoolIDS = append(existingHoerHoolIDS, gtHoerHool.ID)
//		}
//
//		DB.DB.Where("menu_form_id = ? AND id NOT IN ?", galTogoo2.ID, existingHoerHoolIDS).Delete(models.SubMenuHoerGtHoer{})
//
//		existingSalatIDS := []int{}
//		for _, salat := range salats {
//			gtSalat := models.SubMenuSalatGtHoer{}
//
//			DB.DB.Where("food_salat_id = ? AND menu_form_id = ?", salat.FoodSalatID, galTogoo2.ID).Find(&salat)
//
//			if gtSalat.ID <= 0 {
//				gtSalat.MenuFormID = &galTogoo2.ID
//				gtSalat.FoodSalatID = salat.FoodSalatID
//				DB.DB.Create(&gtSalat)
//			}
//
//			existingSalatIDS = append(existingSalatIDS, gtSalat.ID)
//		}
//
//		DB.DB.Where("menu_form_id = ? AND id NOT IN ?", galTogoo2.ID, existingSalatIDS).Delete(models.SubMenuSalatGtHoer{})
//
//	} else {
//		galTogoo2.MainMenuID = &menu.ID
//		galTogoo2.SetDate = menu.SetDate
//		galTogoo2.SetName = menu.SetName
//		galTogoo2.FoodTimeTypeID = menu.FoodTimeTypeID
//
//		DB.DB.Create(&galTogoo2)
//
//		for _, uuhYum := range uuhYums {
//			gtUuhYum := models.SubMenuUuhimGtHoer{}
//
//			gtUuhYum.MenuFormID = &galTogoo2.ID
//			gtUuhYum.FoodUuhimID = uuhYum.FoodUuhimID
//
//			DB.DB.Create(&gtUuhYum)
//		}
//
//		for _, desert := range deserts {
//			gtDesert := models.SubMenuDesertGtHoer{}
//
//			gtDesert.MenuFormID = &galTogoo2.ID
//			gtDesert.FoodDesertID = desert.FoodDesertID
//
//			DB.DB.Create(&gtDesert)
//		}
//
//		for _, negHool := range negHools {
//			gtNegHool := models.SubMenuNegGtHoer{}
//
//			gtNegHool.MenuFormID = &galTogoo2.ID
//			gtNegHool.FoodNegID = negHool.FoodNegID
//
//			DB.DB.Create(&gtNegHool)
//		}
//
//		for _, hoerHool := range hoerHools {
//			gtHoerHool := models.SubMenuHoerGtHoer{}
//
//			gtHoerHool.MenuFormID = &galTogoo2.ID
//			gtHoerHool.FoodHoerID = hoerHool.FoodHoerID
//
//			DB.DB.Create(&gtHoerHool)
//		}
//
//	}
//}
//
//func SaveGalTogoo3(menu *formModels.TblMenu21, uuhYums []models.SubMenuUuhim, deserts []models.SubMenuDesert, negHools []models.SubMenuNeg, hoerHools []models.SubMenuHoer, salats []models.SubMenuSalat) {
//
//	galTogoo3 := models.TblMenuGtGuraw{}
//	DB.DB.Where("main_menu_id = ?", menu.ID).Find(&galTogoo3)
//	if galTogoo3.ID >= 1 {
//		//DO update
//
//		existingUuhYUmsIDS := []int{}
//		for _, uuhYum := range uuhYums {
//			gtUuhYum := models.SubMenuUuhimGtGuraw{}
//
//			DB.DB.Where("food_uuhim_id = ? AND menu_form_id = ?", uuhYum.FoodUuhimID, galTogoo3.ID).Find(&gtUuhYum)
//
//			if gtUuhYum.ID <= 0 {
//				gtUuhYum.MenuFormID = &galTogoo3.ID
//				gtUuhYum.FoodUuhimID = uuhYum.FoodUuhimID
//				DB.DB.Create(&gtUuhYum)
//			}
//
//			existingUuhYUmsIDS = append(existingUuhYUmsIDS, gtUuhYum.ID)
//
//		}
//
//		DB.DB.Where("menu_form_id = ? AND id NOT IN ?", galTogoo3.ID, existingUuhYUmsIDS).Delete(models.SubMenuUuhimGtGuraw{})
//
//		existingDesertIDS := []int{}
//		for _, desert := range deserts {
//			gtDesert := models.SubMenuDesertGtGuraw{}
//
//			DB.DB.Where("food_desert_id = ? AND menu_form_id = ?", desert.FoodDesertID, galTogoo3.ID).Find(&gtDesert)
//
//			if gtDesert.ID <= 0 {
//				gtDesert.MenuFormID = &galTogoo3.ID
//				gtDesert.FoodDesertID = desert.FoodDesertID
//				DB.DB.Create(&gtDesert)
//			}
//			existingDesertIDS = append(existingDesertIDS, gtDesert.ID)
//		}
//
//		DB.DB.Where("menu_form_id = ? AND id NOT IN ?", galTogoo3.ID, existingDesertIDS).Delete(models.SubMenuDesertGtGuraw{})
//
//		existingNegHoolsIDS := []int{}
//		for _, negHool := range negHools {
//			gtNegHool := models.SubMenuNegGtGuraw{}
//
//			DB.DB.Where("food_neg_id = ? AND menu_form_id = ?", negHool.FoodNegID, galTogoo3.ID).Find(&gtNegHool)
//
//			if gtNegHool.ID <= 0 {
//				gtNegHool.MenuFormID = &galTogoo3.ID
//				gtNegHool.FoodNegID = negHool.FoodNegID
//				DB.DB.Create(&gtNegHool)
//			}
//
//			existingNegHoolsIDS = append(existingNegHoolsIDS, gtNegHool.ID)
//
//		}
//
//		DB.DB.Where("menu_form_id = ? AND id NOT IN ?", galTogoo3.ID, existingNegHoolsIDS).Delete(models.SubMenuNegGtGuraw{})
//
//		existingHoerHoolIDS := []int{}
//		for _, hoerHool := range hoerHools {
//			gtHoerHool := models.SubMenuHoerGtGuraw{}
//
//			DB.DB.Where("food_hoer_id = ? AND menu_form_id = ?", hoerHool.FoodHoerID, galTogoo3.ID).Find(&gtHoerHool)
//
//			if gtHoerHool.ID <= 0 {
//				gtHoerHool.MenuFormID = &galTogoo3.ID
//				gtHoerHool.FoodHoerID = hoerHool.FoodHoerID
//				DB.DB.Create(&gtHoerHool)
//			}
//
//			existingHoerHoolIDS = append(existingHoerHoolIDS, gtHoerHool.ID)
//		}
//
//		DB.DB.Where("menu_form_id = ? AND id NOT IN ?", galTogoo3.ID, existingHoerHoolIDS).Delete(models.SubMenuHoerGtGuraw{})
//
//		existingSalatIDS := []int{}
//		for _, salat := range salats {
//			gtSalat := models.SubMenuSalatGtGuraw{}
//
//			DB.DB.Where("food_salat_id = ? AND menu_form_id = ?", salat.FoodSalatID, galTogoo3.ID).Find(&salat)
//
//			if gtSalat.ID <= 0 {
//				gtSalat.MenuFormID = &galTogoo3.ID
//				gtSalat.FoodSalatID = salat.FoodSalatID
//				DB.DB.Create(&gtSalat)
//			}
//
//			existingSalatIDS = append(existingSalatIDS, gtSalat.ID)
//		}
//
//		DB.DB.Where("menu_form_id = ? AND id NOT IN ?", galTogoo3.ID, existingSalatIDS).Delete(models.SubMenuSalatGtGuraw{})
//
//	} else {
//		galTogoo3.MainMenuID = &menu.ID
//		galTogoo3.SetDate = menu.SetDate
//		galTogoo3.SetName = menu.SetName
//		galTogoo3.FoodTimeTypeID = menu.FoodTimeTypeID
//
//		DB.DB.Create(&galTogoo3)
//
//		for _, uuhYum := range uuhYums {
//			gtUuhYum := models.SubMenuUuhimGtGuraw{}
//
//			gtUuhYum.MenuFormID = &galTogoo3.ID
//			gtUuhYum.FoodUuhimID = uuhYum.FoodUuhimID
//
//			DB.DB.Create(&gtUuhYum)
//		}
//
//		for _, desert := range deserts {
//			gtDesert := models.SubMenuDesertGtGuraw{}
//
//			gtDesert.MenuFormID = &galTogoo3.ID
//			gtDesert.FoodDesertID = desert.FoodDesertID
//
//			DB.DB.Create(&gtDesert)
//		}
//
//		for _, negHool := range negHools {
//			gtNegHool := models.SubMenuNegGtGuraw{}
//
//			gtNegHool.MenuFormID = &galTogoo3.ID
//			gtNegHool.FoodNegID = negHool.FoodNegID
//
//			DB.DB.Create(&gtNegHool)
//		}
//
//		for _, hoerHool := range hoerHools {
//			gtHoerHool := models.SubMenuHoerGtGuraw{}
//
//			gtHoerHool.MenuFormID = &galTogoo3.ID
//			gtHoerHool.FoodHoerID = hoerHool.FoodHoerID
//
//			DB.DB.Create(&gtHoerHool)
//		}
//
//	}
//}

func MenuAfterInsertUpdate(menuPre interface{}) {
	menu := menuPre.(*formModels.TblMenu157)
	// subMenu := subMenuPre.(*formModels.SubMenu158)

	// subMenuPre := models.SubMenu{}

	mainSubMenus := []models.SubMenu{}
	mainSubMenuFoods := []models.SubMenuFoods{}

	DB.DB.Where("menu_form_id = ?", menu.ID).Find(&mainSubMenus)
	DB.DB.Where("sub_menu_form_id = ?", menu.ID).Find(&mainSubMenuFoods)

	fmt.Println("Hello", mainSubMenuFoods)
	// fmt.Println("HelloSub", mainSubMenus)

	saveGalTogoo1(menu, mainSubMenus, mainSubMenuFoods)
}

func saveGalTogoo1(menu *formModels.TblMenu157, mainSubMenus []models.SubMenu, mainSubMenuFoods []models.SubMenuFoods) {
	galTogoo1 := models.TblMenuGtNeg{}

	DB.DB.Where("main_menu_id = ?", menu.ID).Find(&galTogoo1)

	if galTogoo1.ID >= 1 {

		//UPDATE

		existingSubMenuIDS := []int{}

		for _, mainSubMenu := range mainSubMenus {
			gtNegSubMenu := models.SubMenuGtNeg{}

			DB.DB.Where("food_type_id = ? AND menu_form_id = ?", gtNegSubMenu.FoodTypeID, galTogoo1.ID).Find(&gtNegSubMenu)

			if gtNegSubMenu.ID <= 0 {
				gtNegSubMenu.MenuFormID = &galTogoo1.ID
				gtNegSubMenu.FoodTypeID = mainSubMenu.FoodTypeID

				DB.DB.Create(&gtNegSubMenu)

				existingSubMenuIDS = append(existingSubMenuIDS, gtNegSubMenu.ID)

				for _, mainSubMenuFood := range mainSubMenuFoods {
					gtNegSubMenuFood := models.SubMenuFoods{}

					DB.DB.Where("food_id = ? AND sub_menu_form_id = ?", gtNegSubMenuFood.FoodID, gtNegSubMenu.ID).Find(&gtNegSubMenuFood)

					if gtNegSubMenuFood.ID <= 0 {
						gtNegSubMenuFood.FoodID = mainSubMenuFood.FoodID
						gtNegSubMenuFood.SubMenuFormID = &gtNegSubMenu.ID

						DB.DB.Create(&gtNegSubMenuFood)
					}
				}
				DB.DB.Where("menu_form_id = ? AND id NOT IN ?", galTogoo1.ID, existingSubMenuIDS).Delete(models.SubMenuGtNeg{})
			}
		}
	} else {

		//INSERT

		galTogoo1.MainMenuID = &menu.ID
		galTogoo1.SetName = menu.SetName
		galTogoo1.SetDate = menu.SetDate
		galTogoo1.OrderRuleID = menu.OrderRuleID

		DB.DB.Create(&galTogoo1)

		for _, mainSubMenu := range mainSubMenus {

			gtNegSubMenu := models.SubMenuGtNeg{}

			fmt.Println(mainSubMenuFoods)

			gtNegSubMenu.MenuFormID = &galTogoo1.ID
			gtNegSubMenu.FoodTypeID = mainSubMenu.FoodTypeID

			DB.DB.Create(&gtNegSubMenu)

			for _, mainSubMenuFood := range mainSubMenuFoods {

				gtNegSubMenuFood := models.SubMenuFoodsGtNeg{}

				fmt.Println("Create", mainSubMenuFood)

				gtNegSubMenuFood.SubMenuFormID = &mainSubMenu.ID
				gtNegSubMenuFood.FoodID = mainSubMenuFood.FoodID

				DB.DB.Create(&gtNegSubMenuFood)
			}
		}
	}
}
