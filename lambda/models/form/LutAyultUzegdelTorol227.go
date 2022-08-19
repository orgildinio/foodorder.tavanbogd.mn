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

func LutAyultUzegdelTorol227Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Аюулт үзэгдлын төрөл",
		Identity: "id",
		Table:    "lut_ayult_uzegdel_torol",
		Model:    new(formModels.LutAyultUzegdelTorol227),
		FieldTypes: map[string]string{
			"id":                  "Text",
			"ayult_uzegdel_torol": "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"ayult_uzegdel_torol": []string{}},
		ValidationMessages: govalidator.MapData{

			"ayult_uzegdel_torol": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
