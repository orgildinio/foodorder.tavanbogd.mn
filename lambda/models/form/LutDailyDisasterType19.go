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

func LutDailyDisasterType19Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "Гамшгийн төрөл",
		Identity: "id",
		Table:    "lut_daily_disaster_type",
		Model:    new(formModels.LutDailyDisasterType19),
		FieldTypes: map[string]string{
			"id":                  "Text",
			"daily_disaster_type": "Text",
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
