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

func LutActivityType15Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Үйл ажиллагааны чиглэл ",
		Identity: "id",
		Table:    "lut_activity_type",
		Model:    new(formModels.LutActivityType15),
		FieldTypes: map[string]string{
			"id":            "Text",
			"activity_type": "Text",
		},
		Formulas:           []models.Formula{},
		ValidationRules:    govalidator.MapData{},
		ValidationMessages: govalidator.MapData{},
		SubForms:           []map[string]interface{}{},
		AfterInsert:        nil,
		AfterUpdate:        nil,
		BeforeInsert:       nil,
		BeforeUpdate:       nil,
		TriggerNameSpace:   "",
	}
}
