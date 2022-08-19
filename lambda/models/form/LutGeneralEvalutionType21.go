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

func LutGeneralEvalutionType21Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Ерөнхий үнэлгээний төрөл org",
		Identity: "id",
		Table:    "lut_general_evalution_type",
		Model:    new(formModels.LutGeneralEvalutionType21),
		FieldTypes: map[string]string{
			"id":                     "Text",
			"general_evalution_type": "Text",
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
