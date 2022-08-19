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

func LutMalTorol171Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Малын төрөл",
		Identity: "id",
		Table:    "lut_mal_torol",
		Model:    new(formModels.LutMalTorol171),
		FieldTypes: map[string]string{
			"id":        "Text",
			"mal_torol": "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"mal_torol": []string{}},
		ValidationMessages: govalidator.MapData{

			"mal_torol": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
