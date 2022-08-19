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

func LutDamageType299Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Хохирлын төрөл ",
		Identity: "id",
		Table:    "lut_damage_type",
		Model:    new(formModels.LutDamageType299),
		FieldTypes: map[string]string{
			"id":          "Text",
			"damage_type": "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"damage_type": []string{}},
		ValidationMessages: govalidator.MapData{

			"damage_type": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
