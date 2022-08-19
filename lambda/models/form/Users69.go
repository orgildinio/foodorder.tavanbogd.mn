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

func Users69Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "ААНБ-д эрх олгох",
		Identity: "id",
		Table:    "users",
		Model:    new(formModels.Users69),
		FieldTypes: map[string]string{
			"id":              "Text",
			"created_at":      "Text",
			"updated_at":      "Text",
			"deleted_at":      "Text",
			"status":          "Text",
			"role":            "Select",
			"login":           "Text",
			"email":           "Text",
			"register_number": "Text",
			"avatar":          "Text",
			"bio":             "Text",
			"first_name":      "Text",
			"last_name":       "Text",
			"birthday":        "Text",
			"phone":           "Text",
			"gender":          "Text",
			"password":        "Password",
			"fcm_token":       "Text",
			"b_id":            "Select",
		},
		Formulas: []models.Formula{},
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
