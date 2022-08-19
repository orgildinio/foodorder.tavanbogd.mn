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

func LutSubReasonHunees113Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Дэд төрөл /Хүнээс/",
		Identity: "id",
		Table:    "lut_sub_reason_hunees",
		Model:    new(formModels.LutSubReasonHunees113),
		FieldTypes: map[string]string{
			"id":                  "Text",
			"shaltgaan_hunees_id": "Select",
			"sub_reason_hunees":   "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"shaltgaan_hunees_id": []string{},
			"sub_reason_hunees":   []string{}},
		ValidationMessages: govalidator.MapData{

			"shaltgaan_hunees_id": []string{},
			"sub_reason_hunees":   []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
