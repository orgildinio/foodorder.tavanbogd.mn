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

func LutHogjilBerhsheelTorol317Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Хөгжлийин бэрхшээлийн төрөл",
		Identity: "id",
		Table:    "lut_hogjil_berhsheel_torol",
		Model:    new(formModels.LutHogjilBerhsheelTorol317),
		FieldTypes: map[string]string{
			"id":                     "Text",
			"hogjil_berhsheel_torol": "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"hogjil_berhsheel_torol": []string{}},
		ValidationMessages: govalidator.MapData{

			"hogjil_berhsheel_torol": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
