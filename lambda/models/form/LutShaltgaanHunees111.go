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

func LutShaltgaanHunees111Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Төрөл /Хүнээс/",
		Identity: "id",
		Table:    "lut_shaltgaan_hunees",
		Model:    new(formModels.LutShaltgaanHunees111),
		FieldTypes: map[string]string{
			"id":                   "Text",
			"lut_shaltgaan_hunees": "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"lut_shaltgaan_hunees": []string{}},
		ValidationMessages: govalidator.MapData{

			"lut_shaltgaan_hunees": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
