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

func LutBiologyDisasterType134Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Биологийн гаралтай аюулт үзэгдлийн төрөл",
		Identity: "id",
		Table:    "lut_biology_disaster_type",
		Model:    new(formModels.LutBiologyDisasterType134),
		FieldTypes: map[string]string{
			"id":                    "Text",
			"biology_disaster_type": "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"biology_disaster_type": []string{}},
		ValidationMessages: govalidator.MapData{

			"biology_disaster_type": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
