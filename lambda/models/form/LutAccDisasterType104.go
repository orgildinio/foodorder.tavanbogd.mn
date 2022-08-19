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

func LutAccDisasterType104Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Ослын төрөл",
		Identity: "id",
		Table:    "lut_acc_disaster_type",
		Model:    new(formModels.LutAccDisasterType104),
		FieldTypes: map[string]string{
			"id":            "Text",
			"disaster_type": "Text",
			"description":   "Textarea",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"disaster_type": []string{}},
		ValidationMessages: govalidator.MapData{

			"disaster_type": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
