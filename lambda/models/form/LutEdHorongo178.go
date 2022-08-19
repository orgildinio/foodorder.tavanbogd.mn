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

func LutEdHorongo178Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Эд хөрөнгө",
		Identity: "id",
		Table:    "lut_ed_horongo",
		Model:    new(formModels.LutEdHorongo178),
		FieldTypes: map[string]string{
			"id":         "Text",
			"ed_horongo": "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"ed_horongo": []string{}},
		ValidationMessages: govalidator.MapData{

			"ed_horongo": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
