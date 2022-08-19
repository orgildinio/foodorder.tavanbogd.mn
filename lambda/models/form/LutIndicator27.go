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

func LutIndicator27Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Индикатор",
		Identity: "id",
		Table:    "lut_indicator",
		Model:    new(formModels.LutIndicator27),
		FieldTypes: map[string]string{
			"id":           "Text",
			"indicator":    "Text",
			"indicator_id": "Select",
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
