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

func LutEvalutionSector11Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Үнэлгээ хийх салбар ",
		Identity: "id",
		Table:    "lut_evalution_sector",
		Model:    new(formModels.LutEvalutionSector11),
		FieldTypes: map[string]string{
			"id":               "Text",
			"evalution_sector": "Text",
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
