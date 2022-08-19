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

func LutMalNasTorol173Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Малын нас",
		Identity: "id",
		Table:    "lut_mal_nas_torol",
		Model:    new(formModels.LutMalNasTorol173),
		FieldTypes: map[string]string{
			"id":            "Text",
			"mal_nas_torol": "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"mal_nas_torol": []string{}},
		ValidationMessages: govalidator.MapData{

			"mal_nas_torol": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
