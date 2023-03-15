package controllers

import (
	"github.com/lambda-platform/lambda/DB"
	"lambda/app/models"
	"lambda/lambda/models/form/formModels"
)

func MenuAfterInsertUpdate(menuPre interface{}) {
	var menu *formModels.TblMenu157 = menuPre.(*formModels.TblMenu157)

	kitchens := []models.Kitchen{}
	DB.DB.Find(&kitchens)

	subMenus := []models.SubMenu{}
	subMenuFoods := []models.SubMenuFoods{}
	subMenuID := models.SubMenu{}
	DB.DB.Find(&subMenuID)

	DB.DB.Where("menu_id = ?", menu.ID).Find(&subMenus)
	DB.DB.Where("sub_menu_id = ?", subMenuID.ID).Find(&subMenuFoods)

	CreateKitchensMenu(menu, subMenus, subMenuFoods)

}

func CreateKitchensMenu(menu *formModels.TblMenu157, subMenus []models.SubMenu, subMenuFoods []models.SubMenuFoods) {

	kitchens := []models.Kitchen{}
	DB.DB.Find(&kitchens)

	for _, kitchen := range kitchens {
		branchKitchen := models.TblMenuGtNeg{}
		DB.DB.Where("main_menu_id = ? AND id = ?", menu.ID, branchKitchen.ID).Find(&branchKitchen)

		branchKitchen.KitchenID = kitchen.ID
		branchKitchen.MainMenuID = menu.ID
		//branchKitchen.OrderRuleID = menu.OrderRuleID
		branchKitchen.SetName = menu.SetName
		//branchKitchen.SetDate = menu.SetDate

		DB.DB.Create(&branchKitchen)

		for _, submenu := range subMenus {
			branchKitchenSubMenu := models.SubMenuGtNeg{}
			DB.DB.Where("id = ? AND menu_id = ?", branchKitchen.ID, branchKitchenSubMenu.MenuID).Find(&branchKitchenSubMenu)

			*branchKitchenSubMenu.FoodTypeID = submenu.FoodTypeID
			branchKitchenSubMenu.MenuID = branchKitchen.ID

			DB.DB.Save(&branchKitchenSubMenu)

			subMenuFoods1 := []models.SubMenuFoods{}
			DB.DB.Where("sub_menu_id = ?", submenu.ID).Find(&subMenuFoods1)

			for _, subMenuFood := range subMenuFoods1 {

				branchKitchenSubMenuFood := models.SubMenuFoodsGtNeg{}
				DB.DB.Where("id = ? AND food_id = ?", branchKitchenSubMenuFood.ID, subMenuFood.FoodID).Find(&branchKitchenSubMenuFood)

				//branchKitchenSubMenuFood.FoodID = subMenuFood.FoodID
				branchKitchenSubMenuFood.SubMenuID = branchKitchenSubMenu.ID

				DB.DB.Debug().Create(&branchKitchenSubMenuFood)

			}
		}
	}
}

//func MenuAfterInsertUpdate(menuPre interface{}) {
//    var menu *formModels.TblMenu157 = menuPre.(*formModels.TblMenu157)
//
//    subMenus := []models.SubMenu{}
//    subMenuFoods := []models.SubMenuFoods{}
//    subMenuID := models.SubMenu{}
//    DB.DB.Find(&subMenuID)
//
//    DB.DB.Where("menu_id = ?", menu.ID).Find(&subMenus)
//    DB.DB.Where("sub_menu_id = ?", subMenuID.ID).Find(&subMenuFoods)
//
//    saveGalTogoo1(menu, subMenus, subMenuFoods)
//    saveGalTogoo2(menu, subMenus, subMenuFoods)
//    saveGalTogoo3(menu, subMenus, subMenuFoods)
//}

//func saveGalTogoo1(menu *formModels.TblMenu157, subMenus []models.SubMenu, subMenuFoods []models.SubMenuFoods) {
//	galTogoo1 := models.TblMenuGtNeg{}
//
//	DB.DB.Where("main_menu_id = ?", menu.ID).Find(&galTogoo1)
//
//	if galTogoo1.ID >= 1 {
//		//Do Update
//
//		//galTogoo1.ID = menu.ID
//		galTogoo1.MainMenuID = &menu.ID
//		galTogoo1.OrderRuleID = menu.OrderRuleID
//		galTogoo1.SetName = menu.SetName
//
//		DB.DB.Save(&galTogoo1)
//
//		existingSubMenuID := []int{}
//
//		for _, subMenu := range subMenus {
//			gt1SubMenu := models.SubMenuGtNeg{}
//
//			DB.DB.Where("food_type_id = ? AND menu_id = ?", subMenu.FoodTypeID, galTogoo1.ID).Find(&gt1SubMenu)
//
//			if gt1SubMenu.ID <= 0 {
//				//gt1SubMenu.ID = subMenu.ID
//				gt1SubMenu.MenuID = &galTogoo1.ID
//				gt1SubMenu.FoodTypeID = subMenu.FoodTypeID
//
//				DB.DB.Create(&gt1SubMenu)
//
//				existingSubMenuFoodID := []int{}
//
//				for _, subMenuFood := range subMenuFoods {
//					gt1SubMenuFood := models.SubMenuFoodsGtNeg{}
//
//					DB.DB.Where("food_id = ? AND sub_menu_id = ?", subMenuFood.FoodID, gt1SubMenu.ID).Find(&gt1SubMenuFood)
//
//					if gt1SubMenuFood.ID <= 0 {
//						gt1SubMenuFood.SubMenuID = &gt1SubMenu.ID
//						gt1SubMenuFood.FoodID = subMenuFood.FoodID
//
//						DB.DB.Save(&gt1SubMenuFood)
//					}
//
//					existingSubMenuFoodID = append(existingSubMenuFoodID, gt1SubMenuFood.ID)
//				}
//
//				DB.DB.Where("sub_menu_id = ? AND id NOT IN ?", gt1SubMenu.ID, existingSubMenuFoodID).Delete(models.SubMenuFoodsGtNeg{})
//			}
//
//			existingSubMenuID = append(existingSubMenuID, gt1SubMenu.ID)
//		}
//
//		DB.DB.Where("menu_id = ? AND id NOT IN ?", galTogoo1.ID, existingSubMenuID).Delete(models.SubMenuGtNeg{})
//
//	} else {
//		galTogoo1.ID = menu.ID
//		galTogoo1.MainMenuID = &menu.ID
//		galTogoo1.OrderRuleID = menu.OrderRuleID
//		galTogoo1.SetName = menu.SetName
//
//		DB.DB.Create(&galTogoo1)
//
//		for _, subMenu := range subMenus {
//			gt1SubMenu := models.SubMenuGtNeg{}
//
//			gt1SubMenu.ID = subMenu.ID
//			gt1SubMenu.MenuID = subMenu.MenuID
//			gt1SubMenu.FoodTypeID = subMenu.FoodTypeID
//
//			DB.DB.Create(&gt1SubMenu)
//
//			for _, subMenuFood := range subMenuFoods {
//				gt1SubMenuFood := models.SubMenuFoodsGtNeg{}
//
//				gt1SubMenuFood.SubMenuID = subMenuFood.SubMenuID
//				gt1SubMenuFood.FoodID = subMenuFood.FoodID
//
//				DB.DB.Create(&gt1SubMenuFood)
//			}
//
//		}
//	}
//}
//
//func saveGalTogoo2(menu *formModels.TblMenu157, subMenus []models.SubMenu, subMenuFoods []models.SubMenuFoods) {
//	galTogoo2 := models.TblMenuGtHoer{}
//
//	DB.DB.Where("main_menu_id = ?", menu.ID).Find(&galTogoo2)
//
//	if galTogoo2.ID >= 1 {
//		//Do Update
//
//		galTogoo2.ID = menu.ID
//		galTogoo2.MainMenuID = &menu.ID
//		galTogoo2.OrderRuleID = menu.OrderRuleID
//		//galTogoo2.SetDate = menu.SetDate
//		galTogoo2.SetName = menu.SetName
//
//		DB.DB.Save(&galTogoo2)
//
//		existingSubMenuID := []int{}
//
//		for _, subMenu := range subMenus {
//			gt2SubMenu := models.SubMenuGtHoer{}
//
//			DB.DB.Where("food_type_id = ? AND menu_id = ?", subMenu.FoodTypeID, galTogoo2.ID).Find(&gt2SubMenu)
//
//			if gt2SubMenu.ID <= 0 {
//				//fmt.Println("hihihi")
//				//gt2SubMenu.ID = subMenu.ID
//				gt2SubMenu.MenuID = &galTogoo2.ID
//				gt2SubMenu.FoodTypeID = subMenu.FoodTypeID
//
//				DB.DB.Create(&gt2SubMenu)
//
//				existingSubMenuFoodID := []int{}
//
//				for _, subMenuFood := range subMenuFoods {
//					gt2SubMenuFood := models.SubMenuFoodsGtHoer{}
//
//					DB.DB.Where("food_id = ? AND sub_menu_id = ?", subMenuFood.FoodID, gt2SubMenu.ID).Find(&gt2SubMenuFood)
//
//					if gt2SubMenuFood.ID <= 0 {
//						gt2SubMenuFood.SubMenuID = &gt2SubMenu.ID
//						gt2SubMenuFood.FoodID = subMenuFood.FoodID
//
//						DB.DB.Save(&gt2SubMenuFood)
//					}
//
//					existingSubMenuFoodID = append(existingSubMenuFoodID, gt2SubMenuFood.ID)
//				}
//
//				DB.DB.Where("sub_menu_id = ? AND id NOT IN ?", gt2SubMenu.ID, existingSubMenuFoodID).Delete(models.SubMenuFoodsGtHoer{})
//			}
//
//			existingSubMenuID = append(existingSubMenuID, gt2SubMenu.ID)
//		}
//
//		DB.DB.Where("menu_id = ? AND id NOT IN ?", galTogoo2.ID, existingSubMenuID).Delete(models.SubMenuGtHoer{})
//
//	} else {
//		galTogoo2.ID = menu.ID
//		galTogoo2.MainMenuID = &menu.ID
//		galTogoo2.OrderRuleID = menu.OrderRuleID
//		//galTogoo2.SetDate = menu.SetDate
//		galTogoo2.SetName = menu.SetName
//
//		DB.DB.Create(&galTogoo2)
//
//		for _, subMenu := range subMenus {
//			gt2SubMenu := models.SubMenuGtHoer{}
//
//			gt2SubMenu.ID = subMenu.ID
//			gt2SubMenu.MenuID = subMenu.MenuID
//			gt2SubMenu.FoodTypeID = subMenu.FoodTypeID
//
//			DB.DB.Create(&gt2SubMenu)
//
//			for _, subMenuFood := range subMenuFoods {
//				gt2SubMenuFood := models.SubMenuFoodsGtHoer{}
//
//				gt2SubMenuFood.SubMenuID = subMenuFood.SubMenuID
//				gt2SubMenuFood.FoodID = subMenuFood.FoodID
//
//				DB.DB.Create(&gt2SubMenuFood)
//			}
//
//		}
//	}
//
//}
//
//func saveGalTogoo3(menu *formModels.TblMenu157, subMenus []models.SubMenu, subMenuFoods []models.SubMenuFoods) {
//	galTogoo3 := models.TblMenuGtGuraw{}
//
//	DB.DB.Where("main_menu_id = ?", menu.ID).Find(&galTogoo3)
//
//	if galTogoo3.ID >= 1 {
//		//Do Update
//
//		galTogoo3.ID = menu.ID
//		galTogoo3.MainMenuID = &menu.ID
//		galTogoo3.OrderRuleID = menu.OrderRuleID
//		//galTogoo3.SetDate = menu.SetDate
//		galTogoo3.SetName = menu.SetName
//
//		DB.DB.Save(&galTogoo3)
//
//		existingSubMenuID := []int{}
//
//		for _, subMenu := range subMenus {
//			gt3SubMenu := models.SubMenuGtGuraw{}
//
//			DB.DB.Where("food_type_id = ? AND menu_id = ?", subMenu.FoodTypeID, galTogoo3.ID).Find(&gt3SubMenu)
//
//			if gt3SubMenu.ID <= 0 {
//				//fmt.Println("hihihi")
//				//gt3SubMenu.ID = subMenu.ID
//				gt3SubMenu.MenuID = &galTogoo3.ID
//				gt3SubMenu.FoodTypeID = subMenu.FoodTypeID
//
//				DB.DB.Create(&gt3SubMenu)
//
//				existingSubMenuFoodID := []int{}
//
//				for _, subMenuFood := range subMenuFoods {
//					gt3SubMenuFood := models.SubMenuFoodsGtGuraw{}
//
//					DB.DB.Where("food_id = ? AND sub_menu_id = ?", subMenuFood.FoodID, gt3SubMenu.ID).Find(&gt3SubMenuFood)
//
//					if gt3SubMenuFood.ID <= 0 {
//						gt3SubMenuFood.SubMenuID = &gt3SubMenu.ID
//						gt3SubMenuFood.FoodID = subMenuFood.FoodID
//
//						DB.DB.Save(&gt3SubMenuFood)
//					}
//
//					existingSubMenuFoodID = append(existingSubMenuFoodID, gt3SubMenuFood.ID)
//				}
//
//				DB.DB.Where("sub_menu_id = ? AND id NOT IN ?", gt3SubMenu.ID, existingSubMenuFoodID).Delete(models.SubMenuFoodsGtGuraw{})
//			}
//
//			existingSubMenuID = append(existingSubMenuID, gt3SubMenu.ID)
//		}
//
//		DB.DB.Where("menu_id = ? AND id NOT IN ?", galTogoo3.ID, existingSubMenuID).Delete(models.SubMenuGtGuraw{})
//
//	} else {
//		galTogoo3.ID = menu.ID
//		galTogoo3.MainMenuID = &menu.ID
//		galTogoo3.OrderRuleID = menu.OrderRuleID
//		//galTogoo3.SetDate = menu.SetDate
//		galTogoo3.SetName = menu.SetName
//
//		DB.DB.Create(&galTogoo3)
//
//		for _, subMenu := range subMenus {
//			gt3SubMenu := models.SubMenuGtGuraw{}
//
//			gt3SubMenu.ID = subMenu.ID
//			gt3SubMenu.MenuID = subMenu.MenuID
//			gt3SubMenu.FoodTypeID = subMenu.FoodTypeID
//
//			DB.DB.Create(&gt3SubMenu)
//
//			for _, subMenuFood := range subMenuFoods {
//				gt3SubMenuFood := models.SubMenuFoodsGtGuraw{}
//
//				gt3SubMenuFood.SubMenuID = subMenuFood.SubMenuID
//				gt3SubMenuFood.FoodID = subMenuFood.FoodID
//
//				DB.DB.Create(&gt3SubMenuFood)
//			}
//
//		}
//	}
//
//}
