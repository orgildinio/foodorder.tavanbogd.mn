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

func LutSubIndicator29Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Дэд индикатор",
		Identity: "id",
		Table:    "lut_sub_indicator",
		Model:    new(formModels.LutSubIndicator29),
		FieldTypes: map[string]string{
			"id":            "Text",
			"indicator_id":  "Select",
			"sub_indicator": "Text",
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
