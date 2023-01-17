package controllers

import (
	"github.com/lambda-platform/lambda/DB"
	"lambda/app/models"
	"lambda/lambda/models/form/formModels"
)

func MenuAfterInsertUpdate(menuPre interface{}) {
	var menu *formModels.TblMenu157 = menuPre.(*formModels.TblMenu157)

	subMenus := []models.SubMenu{}
	subMenuFoods := []models.SubMenuFoods{}
	subMenuID := models.SubMenu{}
	DB.DB.Find(&subMenuID)

	DB.DB.Where("menu_id = ?", menu.ID).Find(&subMenus)
	DB.DB.Where("sub_menu_id = ?", subMenuID.ID).Find(&subMenuFoods)

	saveGalTogoo1(menu, subMenus, subMenuFoods)
	saveGalTogoo2(menu, subMenus, subMenuFoods)
	saveGalTogoo3(menu, subMenus, subMenuFoods)
}

func saveGalTogoo1(menu *formModels.TblMenu157, subMenus []models.SubMenu, subMenuFoods []models.SubMenuFoods) {
	galTogoo1 := models.TblMenuGtNeg{}

	DB.DB.Where("main_menu_id = ?", menu.ID).Find(&galTogoo1)

	if galTogoo1.ID >= 1 {
		//Do Update

		//galTogoo1.ID = menu.ID
		galTogoo1.MainMenuID = &menu.ID
		galTogoo1.OrderRuleID = menu.OrderRuleID
		galTogoo1.SetDate = menu.SetDate
		galTogoo1.SetName = menu.SetName

		DB.DB.Save(&galTogoo1)

		existingSubMenuID := []int{}

		for _, subMenu := range subMenus {
			gt1SubMenu := models.SubMenuGtNeg{}

			DB.DB.Where("food_type_id = ? AND menu_id = ?", subMenu.FoodTypeID, galTogoo1.ID).Find(&gt1SubMenu)

			if gt1SubMenu.ID <= 0 {
				//gt1SubMenu.ID = subMenu.ID
				gt1SubMenu.MenuID = &galTogoo1.ID
				gt1SubMenu.FoodTypeID = subMenu.FoodTypeID

				DB.DB.Create(&gt1SubMenu)

				existingSubMenuFoodID := []int{}

				for _, subMenuFood := range subMenuFoods {
					gt1SubMenuFood := models.SubMenuFoodsGtNeg{}

					DB.DB.Where("food_id = ? AND sub_menu_id = ?", subMenuFood.FoodID, gt1SubMenu.ID).Find(&gt1SubMenuFood)

					if gt1SubMenuFood.ID <= 0 {
						gt1SubMenuFood.SubMenuID = &gt1SubMenu.ID
						gt1SubMenuFood.FoodID = subMenuFood.FoodID

						DB.DB.Save(&gt1SubMenuFood)
					}

					existingSubMenuFoodID = append(existingSubMenuFoodID, gt1SubMenuFood.ID)
				}

				DB.DB.Where("sub_menu_id = ? AND id NOT IN ?", gt1SubMenu.ID, existingSubMenuFoodID).Delete(models.SubMenuFoodsGtNeg{})
			}

			existingSubMenuID = append(existingSubMenuID, gt1SubMenu.ID)
		}

		DB.DB.Where("menu_id = ? AND id NOT IN ?", galTogoo1.ID, existingSubMenuID).Delete(models.SubMenuGtNeg{})

	} else {
		galTogoo1.ID = menu.ID
		galTogoo1.MainMenuID = &menu.ID
		galTogoo1.OrderRuleID = menu.OrderRuleID
		galTogoo1.SetDate = menu.SetDate
		galTogoo1.SetName = menu.SetName

		DB.DB.Create(&galTogoo1)

		for _, subMenu := range subMenus {
			gt1SubMenu := models.SubMenuGtNeg{}

			gt1SubMenu.ID = subMenu.ID
			gt1SubMenu.MenuID = subMenu.MenuID
			gt1SubMenu.FoodTypeID = subMenu.FoodTypeID

			DB.DB.Create(&gt1SubMenu)

			for _, subMenuFood := range subMenuFoods {
				gt1SubMenuFood := models.SubMenuFoodsGtNeg{}

				gt1SubMenuFood.SubMenuID = subMenuFood.SubMenuID
				gt1SubMenuFood.FoodID = subMenuFood.FoodID

				DB.DB.Create(&gt1SubMenuFood)
			}

		}
	}
}

func saveGalTogoo2(menu *formModels.TblMenu157, subMenus []models.SubMenu, subMenuFoods []models.SubMenuFoods) {
	galTogoo2 := models.TblMenuGtHoer{}

	DB.DB.Where("main_menu_id = ?", menu.ID).Find(&galTogoo2)

	if galTogoo2.ID >= 1 {
		//Do Update

		galTogoo2.ID = menu.ID
		galTogoo2.MainMenuID = &menu.ID
		galTogoo2.OrderRuleID = menu.OrderRuleID
		//galTogoo2.SetDate = menu.SetDate
		galTogoo2.SetName = menu.SetName

		DB.DB.Save(&galTogoo2)

		existingSubMenuID := []int{}

		for _, subMenu := range subMenus {
			gt2SubMenu := models.SubMenuGtHoer{}

			DB.DB.Where("food_type_id = ? AND menu_id = ?", subMenu.FoodTypeID, galTogoo2.ID).Find(&gt2SubMenu)

			if gt2SubMenu.ID <= 0 {
				//fmt.Println("hihihi")
				//gt2SubMenu.ID = subMenu.ID
				gt2SubMenu.MenuID = &galTogoo2.ID
				gt2SubMenu.FoodTypeID = subMenu.FoodTypeID

				DB.DB.Create(&gt2SubMenu)

				existingSubMenuFoodID := []int{}

				for _, subMenuFood := range subMenuFoods {
					gt2SubMenuFood := models.SubMenuFoodsGtHoer{}

					DB.DB.Where("food_id = ? AND sub_menu_id = ?", subMenuFood.FoodID, gt2SubMenu.ID).Find(&gt2SubMenuFood)

					if gt2SubMenuFood.ID <= 0 {
						gt2SubMenuFood.SubMenuID = &gt2SubMenu.ID
						gt2SubMenuFood.FoodID = subMenuFood.FoodID

						DB.DB.Save(&gt2SubMenuFood)
					}

					existingSubMenuFoodID = append(existingSubMenuFoodID, gt2SubMenuFood.ID)
				}

				DB.DB.Where("sub_menu_id = ? AND id NOT IN ?", gt2SubMenu.ID, existingSubMenuFoodID).Delete(models.SubMenuFoodsGtHoer{})
			}

			existingSubMenuID = append(existingSubMenuID, gt2SubMenu.ID)
		}

		DB.DB.Where("menu_id = ? AND id NOT IN ?", galTogoo2.ID, existingSubMenuID).Delete(models.SubMenuGtHoer{})

	} else {
		galTogoo2.ID = menu.ID
		galTogoo2.MainMenuID = &menu.ID
		galTogoo2.OrderRuleID = menu.OrderRuleID
		//galTogoo2.SetDate = menu.SetDate
		galTogoo2.SetName = menu.SetName

		DB.DB.Create(&galTogoo2)

		for _, subMenu := range subMenus {
			gt2SubMenu := models.SubMenuGtHoer{}

			gt2SubMenu.ID = subMenu.ID
			gt2SubMenu.MenuID = subMenu.MenuID
			gt2SubMenu.FoodTypeID = subMenu.FoodTypeID

			DB.DB.Create(&gt2SubMenu)

			for _, subMenuFood := range subMenuFoods {
				gt2SubMenuFood := models.SubMenuFoodsGtHoer{}

				gt2SubMenuFood.SubMenuID = subMenuFood.SubMenuID
				gt2SubMenuFood.FoodID = subMenuFood.FoodID

				DB.DB.Create(&gt2SubMenuFood)
			}

		}
	}

}

func saveGalTogoo3(menu *formModels.TblMenu157, subMenus []models.SubMenu, subMenuFoods []models.SubMenuFoods) {
	galTogoo3 := models.TblMenuGtGuraw{}

	DB.DB.Where("main_menu_id = ?", menu.ID).Find(&galTogoo3)

	if galTogoo3.ID >= 1 {
		//Do Update

		galTogoo3.ID = menu.ID
		galTogoo3.MainMenuID = &menu.ID
		galTogoo3.OrderRuleID = menu.OrderRuleID
		//galTogoo3.SetDate = menu.SetDate
		galTogoo3.SetName = menu.SetName

		DB.DB.Save(&galTogoo3)

		existingSubMenuID := []int{}

		for _, subMenu := range subMenus {
			gt3SubMenu := models.SubMenuGtGuraw{}

			DB.DB.Where("food_type_id = ? AND menu_id = ?", subMenu.FoodTypeID, galTogoo3.ID).Find(&gt3SubMenu)

			if gt3SubMenu.ID <= 0 {
				//fmt.Println("hihihi")
				//gt3SubMenu.ID = subMenu.ID
				gt3SubMenu.MenuID = &galTogoo3.ID
				gt3SubMenu.FoodTypeID = subMenu.FoodTypeID

				DB.DB.Create(&gt3SubMenu)

				existingSubMenuFoodID := []int{}

				for _, subMenuFood := range subMenuFoods {
					gt3SubMenuFood := models.SubMenuFoodsGtGuraw{}

					DB.DB.Where("food_id = ? AND sub_menu_id = ?", subMenuFood.FoodID, gt3SubMenu.ID).Find(&gt3SubMenuFood)

					if gt3SubMenuFood.ID <= 0 {
						gt3SubMenuFood.SubMenuID = &gt3SubMenu.ID
						gt3SubMenuFood.FoodID = subMenuFood.FoodID

						DB.DB.Save(&gt3SubMenuFood)
					}

					existingSubMenuFoodID = append(existingSubMenuFoodID, gt3SubMenuFood.ID)
				}

				DB.DB.Where("sub_menu_id = ? AND id NOT IN ?", gt3SubMenu.ID, existingSubMenuFoodID).Delete(models.SubMenuFoodsGtGuraw{})
			}

			existingSubMenuID = append(existingSubMenuID, gt3SubMenu.ID)
		}

		DB.DB.Where("menu_id = ? AND id NOT IN ?", galTogoo3.ID, existingSubMenuID).Delete(models.SubMenuGtGuraw{})

	} else {
		galTogoo3.ID = menu.ID
		galTogoo3.MainMenuID = &menu.ID
		galTogoo3.OrderRuleID = menu.OrderRuleID
		//galTogoo3.SetDate = menu.SetDate
		galTogoo3.SetName = menu.SetName

		DB.DB.Create(&galTogoo3)

		for _, subMenu := range subMenus {
			gt3SubMenu := models.SubMenuGtGuraw{}

			gt3SubMenu.ID = subMenu.ID
			gt3SubMenu.MenuID = subMenu.MenuID
			gt3SubMenu.FoodTypeID = subMenu.FoodTypeID

			DB.DB.Create(&gt3SubMenu)

			for _, subMenuFood := range subMenuFoods {
				gt3SubMenuFood := models.SubMenuFoodsGtGuraw{}

				gt3SubMenuFood.SubMenuID = subMenuFood.SubMenuID
				gt3SubMenuFood.FoodID = subMenuFood.FoodID

				DB.DB.Create(&gt3SubMenuFood)
			}

		}
	}

}

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

//func MenuAfterInsertUpdate(menuPre interface{}) {
//	menu := menuPre.(*formModels.TblMenu157)
//
//	mainSubMenus := []models.SubMenu{}
//	subMenuFoods := []models.SubMenuFoods{}
//	subMenuID := models.SubMenu{}
//	DB.DB.Find(&subMenuID)
//	//fmt.Println("Testing", subMenuID.ID, *subMenuID.FoodTypeID)
//
//	DB.DB.Where("main_menu_id = ?", menu.ID).Find(&mainSubMenus)
//	DB.DB.Where("sub_menu_id = ?", subMenuID.ID).Find(&subMenuFoods)
//
//	saveGalTogoo1(menu, mainSubMenus, subMenuFoods)
//}
//
//func saveGalTogoo1(menu *formModels.TblMenu157, mainSubMenus []models.SubMenu, subMenuFoods []models.SubMenuFoods) {
//	galTogoo1 := models.TblMenuGtNeg{}
//	DB.DB.Where("main_menu_id = ?", menu.ID).Find(&galTogoo1)
//
//	DB.DB.First(&galTogoo1)
//
//	galTogoo1.ID = menu.ID
//	galTogoo1.MainMenuID = &menu.ID
//	galTogoo1.SetName = menu.SetName
//	//galTogoo1.SetDate = menu.SetDate
//	galTogoo1.OrderRuleID = menu.OrderRuleID
//
//	DB.DB.Save(&galTogoo1)
//
//	if galTogoo1.ID >= 1 {
//		//UPDATE
//
//		existSubMenuID := []int{}
//		for _, mainSubMenu := range mainSubMenus {
//			gtSubMenu := models.SubMenuGtNeg{}
//
//			DB.DB.Where("food_type_id = ? AND menu_id = ?", mainSubMenu.FoodTypeID, galTogoo1.ID).Find(&gtSubMenu)
//
//			existSubMenuFoodID := []int{}
//			if gtSubMenu.ID <= 0 {
//
//				gtSubMenu.ID = mainSubMenu.ID
//				gtSubMenu.MenuID = &galTogoo1.ID
//				gtSubMenu.FoodTypeID = mainSubMenu.FoodTypeID
//
//				DB.DB.Create(&gtSubMenu)
//
//				for _, subMenuFood := range subMenuFoods {
//					gtSubMenuFoods := models.SubMenuFoodsGtNeg{}
//
//					DB.DB.Where("food_id = ? AND sub_menu_id = ?", subMenuFood.FoodID, gtSubMenu.ID).Find(&gtSubMenuFoods)
//
//					if gtSubMenuFoods.ID <= 0 {
//						gtSubMenuFoods.SubMenuID = &mainSubMenu.ID
//						gtSubMenuFoods.FoodID = subMenuFood.FoodID
//
//						DB.DB.Create(&gtSubMenuFoods)
//					}
//
//					existSubMenuFoodID = append(existSubMenuFoodID, gtSubMenuFoods.ID)
//				}
//
//				DB.DB.Where("sub_menu_id = ? AND id NOT IN ?", gtSubMenu.ID, existSubMenuFoodID).Delete(models.SubMenuFoodsGtNeg{})
//			}
//
//			existSubMenuID = append(existSubMenuID, gtSubMenu.ID)
//		}
//
//		DB.DB.Where("menu_id = ? AND id NOT IN ?", galTogoo1.ID, existSubMenuID).Delete(models.SubMenuGtNeg{})
//	} else {
//		galTogoo1.ID = menu.ID
//		galTogoo1.MainMenuID = &menu.ID
//		galTogoo1.SetName = menu.SetName
//		//galTogoo1.SetDate = menu.SetDate
//		galTogoo1.OrderRuleID = menu.OrderRuleID
//
//		DB.DB.Create(&galTogoo1)
//
//		for _, mainSubMenu := range mainSubMenus {
//			gtSubMenu := models.SubMenuGtNeg{}
//
