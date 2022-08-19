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

func LutMaindisasterType84Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "Аюулын төрөл",
		Identity: "id",
		Table:    "lut_maindisaster_type",
		Model:    new(formModels.LutMaindisasterType84),
		FieldTypes: map[string]string{
			"id":                "Text",
			"maindisaster_type": "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"maindisaster_type": []string{}},
		ValidationMessages: govalidator.MapData{

			"maindisaster_type": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
