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

func LutMethodolgyType40Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Аргачлал",
		Identity: "id",
		Table:    "lut_methodolgy_type",
		Model:    new(formModels.LutMethodolgyType40),
		FieldTypes: map[string]string{
			"id":        "Text",
			"argachlal": "Text",
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
