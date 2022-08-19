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

func LutSubSubSubDamageType314Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Хохирлын дэд дэд дэд төрөл ",
		Identity: "id",
		Table:    "lut_sub_sub_sub_damage_type",
		Model:    new(formModels.LutSubSubSubDamageType314),
		FieldTypes: map[string]string{
			"id":                      "Text",
			"sub_sub_damage_type_id":  "Select",
			"sub_sub_sub_damage_type": "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"sub_sub_damage_type_id":  []string{},
			"sub_sub_sub_damage_type": []string{}},
		ValidationMessages: govalidator.MapData{

			"sub_sub_damage_type_id":  []string{},
			"sub_sub_sub_damage_type": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
