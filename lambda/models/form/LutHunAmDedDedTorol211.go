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

func LutHunAmDedDedTorol211Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Хүн амын дэд дэд төрөл ",
		Identity: "id",
		Table:    "lut_hun_am_ded_ded_torol",
		Model:    new(formModels.LutHunAmDedDedTorol211),
		FieldTypes: map[string]string{
			"id":                   "Text",
			"hun_am_ded_torol_id":  "Select",
			"hun_am_ded_ded_torol": "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"hun_am_ded_torol_id":  []string{},
			"hun_am_ded_ded_torol": []string{}},
		ValidationMessages: govalidator.MapData{

			"hun_am_ded_torol_id":  []string{},
			"hun_am_ded_ded_torol": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
