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

func LutMalDedTorol172Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Малын дэд төрөл",
		Identity: "id",
		Table:    "lut_mal_ded_torol",
		Model:    new(formModels.LutMalDedTorol172),
		FieldTypes: map[string]string{
			"id":            "Text",
			"mal_torol_id":  "Select",
			"mal_ded_torol": "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"mal_torol_id":  []string{},
			"mal_ded_torol": []string{}},
		ValidationMessages: govalidator.MapData{

			"mal_torol_id":  []string{},
			"mal_ded_torol": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
