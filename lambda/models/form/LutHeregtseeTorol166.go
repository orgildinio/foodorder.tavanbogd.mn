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

func LutHeregtseeTorol166Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Хэрэгцээний төрөл",
		Identity: "id",
		Table:    "lut_heregtsee_torol",
		Model:    new(formModels.LutHeregtseeTorol166),
		FieldTypes: map[string]string{
			"id":              "Text",
			"heregtsee_torol": "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"heregtsee_torol": []string{}},
		ValidationMessages: govalidator.MapData{

			"heregtsee_torol": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
