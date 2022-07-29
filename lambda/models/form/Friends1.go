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

func Friends1Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "friends",
		Identity: "id",
		Table:    "friends",
		Model:    new(formModels.Friends1),
		FieldTypes: map[string]string{
			"id":    "Text",
			"name":  "Text",
			"phone": "Text",
			"zurag": "ImageDrag",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"name":  []string{"required"},
			"phone": []string{"required"}},
		ValidationMessages: govalidator.MapData{

			"name":  []string{},
			"phone": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
