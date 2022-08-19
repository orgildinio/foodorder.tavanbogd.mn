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

func LutSubAccDisasterType108Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Ослын дэд төрөл",
		Identity: "id",
		Table:    "lut_sub_acc_disaster_type",
		Model:    new(formModels.LutSubAccDisasterType108),
		FieldTypes: map[string]string{
			"id":                    "Text",
			"acc_disaster_type_id":  "Select",
			"sub_acc_disaster_type": "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"acc_disaster_type_id":  []string{},
			"sub_acc_disaster_type": []string{}},
		ValidationMessages: govalidator.MapData{

			"acc_disaster_type_id":  []string{},
			"sub_acc_disaster_type": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
