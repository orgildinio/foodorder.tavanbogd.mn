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

func LutSubReason129Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Халдварласан дэд үзүүлэлт",
		Identity: "id",
		Table:    "lut_sub_reason",
		Model:    new(formModels.LutSubReason129),
		FieldTypes: map[string]string{
			"id":         "",
			"reason_id":  "Select",
			"sub_reason": "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"sub_reason": []string{}},
		ValidationMessages: govalidator.MapData{

			"sub_reason": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
