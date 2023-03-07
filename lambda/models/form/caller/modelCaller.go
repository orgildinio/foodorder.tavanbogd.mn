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

	case "109":
		return form.FoodOrder109Dataform()

	case "86":
		return form.FoodOrder86Dataform()

	case "82":
		return form.FoodOrder82Dataform()

	case "66":
		return form.FoodOrder66Dataform()

	case "8":
		return form.LutOrderType8Dataform()

	case "10":
		return form.LutStaffType10Dataform()

	case "108":
		return form.FoodOrder108Dataform()

	case "98":
		return form.LutMaterialUnit98Dataform()

	case "110":
		return form.FoodOrder110Dataform()

	case "111":
		return form.FoodOrder111Dataform()

	case "14":
		return form.TblCorp14Dataform()

	case "68":
		return form.FoodOrder68Dataform()

	case "102":
		return form.FoodOrderInBasket102Dataform()

	case "17":
		return form.Users17Dataform()

	case "18":
		return form.TblFood18Dataform()

	case "100":
		return form.LutMaterialStatus100Dataform()

	case "164":
		return form.SubMenu164Dataform()

	case "112":
		return form.LutBank112Dataform()

	case "33":
		return form.LutFoodType33Dataform()

	case "114":
		return form.Users114Dataform()

	case "236":
		return form.SubMenuGtNeg236Dataform()

	case "35":
		return form.TblOrderRule35Dataform()

	case "237":
		return form.TblMenuGtNeg237Dataform()

	case "104":
		return form.FoodPrice104Dataform()

	case "53":
		return form.TblMaterial53Dataform()

	case "56":
		return form.LutFoodTimeType56Dataform()

	case "192":
		return form.Orders192Dataform()

	case "117":
		return form.FoodOrder117Dataform()

	case "157":
		return form.TblMenu157Dataform()

	case "63":
		return form.LutKitchenType63Dataform()

	case "70":
		return form.AddLocation70Dataform()

	case "60":
		return form.FoodOrder60Dataform()

	case "199":
		return form.SubOrderFoodType199Dataform()

	case "221":
		return form.Kitchen221Dataform()

	case "189":
		return form.SubMenuGtNeg189Dataform()

	case "229":
		return form.TblMenuGtNeg229Dataform()

	}
	return dataform.Dataform{}

}
