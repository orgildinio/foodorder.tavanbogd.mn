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

func LutHunAmTorol208Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Хүн амын төрөл ",
		Identity: "id",
		Table:    "lut_hun_am_torol",
		Model:    new(formModels.LutHunAmTorol208),
		FieldTypes: map[string]string{
			"id":           "Text",
			"hun_am_torol": "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"hun_am_torol": []string{}},
		ValidationMessages: govalidator.MapData{

			"hun_am_torol": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
