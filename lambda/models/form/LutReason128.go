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

func LutReason128Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Халдварласан үзүүлэлт",
		Identity: "id",
		Table:    "lut_reason",
		Model:    new(formModels.LutReason128),
		FieldTypes: map[string]string{
			"id":     "Text",
			"reason": "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"reason": []string{}},
		ValidationMessages: govalidator.MapData{

			"reason": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
