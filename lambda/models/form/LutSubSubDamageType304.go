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

func LutSubSubDamageType304Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Хохирлын дэд дэд төрөл ",
		Identity: "id",
		Table:    "lut_sub_sub_damage_type",
		Model:    new(formModels.LutSubSubDamageType304),
		FieldTypes: map[string]string{
			"id":                 "Text",
			"damage_sub_type_id": "Select",
			"subsub_damage_type": "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"damage_sub_type_id": []string{},
			"subsub_damage_type": []string{}},
		ValidationMessages: govalidator.MapData{

			"damage_sub_type_id": []string{},
			"subsub_damage_type": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
