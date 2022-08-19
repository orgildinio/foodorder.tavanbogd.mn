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

func LutMeasure170Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Хэмжих нэгж",
		Identity: "id",
		Table:    "lut_measure",
		Model:    new(formModels.LutMeasure170),
		FieldTypes: map[string]string{
			"id":      "Text",
			"measure": "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"measure": []string{}},
		ValidationMessages: govalidator.MapData{

			"measure": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
