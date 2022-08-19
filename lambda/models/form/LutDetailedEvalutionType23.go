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

func LutDetailedEvalutionType23Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "Ерөнхий үнэлгээний төрөл /нэгж/",
		Identity: "id",
		Table:    "lut_detailed_evalution_type",
		Model:    new(formModels.LutDetailedEvalutionType23),
		FieldTypes: map[string]string{
			"id":                      "Text",
			"detailed_evalution_type": "Text",
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
