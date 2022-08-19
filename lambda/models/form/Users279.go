package form

import (
	"github.com/lambda-platform/lambda/DB"
	"github.com/lambda-platform/lambda/dataform"
	"github.com/lambda-platform/lambda/models"
	"github.com/thedevsaddam/govalidator"
	"lambda/lambda/models/form/formModels"
	"time"
)

var _ = time.Time{}
var _ = DB.Date{}

func Users279Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "Хэрэглэгч",
		Identity: "id",
		Table:    "users",
		Model:    new(formModels.Users279),
		FieldTypes: map[string]string{
			"id":               "Text",
			"created_at":       "Text",
			"updated_at":       "Text",
			"deleted_at":       "Text",
			"status":           "Text",
			"role":             "Select",
			"login":            "Text",
			"email":            "Text",
			"register_number":  "Text",
			"avatar":           "Text",
			"bio":              "Text",
			"first_name":       "Text",
			"last_name":        "Text",
			"birthday":         "Date",
			"phone":            "Text",
			"gender":           "Radio",
			"password":         "Password",
			"fcm_token":        "Text",
			"b_id":             "Select",
			"user_type_id":     "Select",
			"alban_tushaal_id": "Select",
			"aimag_id":         "Select",
			"sum_id":           "Select",
			"bag_id":           "Select",
			"angi_id":          "Select",
			"salbari_id":       "Select",
		},
		Formulas: []models.Formula{
			models.Formula{
				Targets: []models.Target{
					models.Target{
						Field: "email",
						Prop:  "hidden",
					}, models.Target{
						Field: "register_number",
						Prop:  "hidden",
					}, models.Target{
						Field: "first_name",
						Prop:  "hidden",
					}, models.Target{
						Field: "last_name",
						Prop:  "hidden",
					}, models.Target{
						Field: "birthday",
						Prop:  "hidden",
					}, models.Target{
						Field: "phone",
						Prop:  "hidden",
					}, models.Target{
						Field: "gender",
						Prop:  "hidden",
					}, models.Target{
						Field: "alban_tushaal_id",
						Prop:  "hidden",
					}, models.Target{
						Field: "aimag_id",
						Prop:  "hidden",
					}, models.Target{
						Field: "bag_id",
						Prop:  "hidden",
					}, models.Target{
						Field: "sum_id",
						Prop:  "hidden",
					}, models.Target{
						Field: "angi_id",
						Prop:  "hidden",
					}, models.Target{
						Field: "salbari_id",
						Prop:  "hidden",
					},
				},
				Template: "'{user_type_id}' != '1'",
				Model:    "",
				Form:     "main",
			}, models.Formula{
				Targets: []models.Target{
					models.Target{
						Field: "b_id",
						Prop:  "hidden",
					},
				},
				Template: "'{user_type_id}' != '2'",
				Model:    "",
				Form:     "main",
			},
		},
		ValidationRules: govalidator.MapData{

			"login":    []string{"required"},
			"password": []string{"required"}},
		ValidationMessages: govalidator.MapData{

			"login":    []string{},
			"password": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
