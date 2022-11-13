package caller

import "lambda/lambda/models/form"
import "github.com/lambda-platform/lambda/dataform"

func GetMODEL(schema_id string) dataform.Dataform {

	switch schema_id {

	case "crud_form":
		return form.KrudDataform()

	case "notification_target_form":
		return form.NotificationTargetsDataform()

	case "menu_form":
		return form.MenuFormDataform()

	case "user_form":
		return form.UserFormDataform()

	case "user_profile":
		return form.UserProfile()

	case "user_password":
		return form.UsersDataform()

	case "1":
		return form.LutBarimtType1Dataform()

	case "4":
		return form.LutPositionType4Dataform()

	case "6":
		return form.LutPaymentType6Dataform()

	case "66":
		return form.FoodOrder66Dataform()

	case "84":
		return form.FoodOrder84Dataform()

	case "86":
		return form.FoodOrder86Dataform()

	case "82":
		return form.FoodOrder82Dataform()

	case "8":
		return form.LutOrderType8Dataform()

	case "10":
		return form.LutStaffType10Dataform()

	case "14":
		return form.TblCorp14Dataform()

	case "83":
		return form.FoodOrder83Dataform()

	case "85":
		return form.FoodOrder85Dataform()

	case "98":
		return form.LutMaterialUnit98Dataform()

	case "68":
		return form.FoodOrder68Dataform()

	case "102":
		return form.FoodOrderInBasket102Dataform()

	case "17":
		return form.Users17Dataform()

	case "100":
		return form.LutMaterialStatus100Dataform()

	case "18":
		return form.TblFood18Dataform()

	case "70":
		return form.AddLocation70Dataform()

	case "21":
		return form.TblMenu21Dataform()

	case "33":
		return form.LutFoodType33Dataform()

	case "35":
		return form.TblOrderRule35Dataform()

	case "53":
		return form.TblMaterial53Dataform()

	case "56":
		return form.LutFoodTimeType56Dataform()

	case "63":
		return form.LutKitchenType63Dataform()

	case "60":
		return form.FoodOrder60Dataform()

	}
	return dataform.Dataform{}

}
