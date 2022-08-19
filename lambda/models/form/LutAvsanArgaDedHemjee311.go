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

func LutAvsanArgaDedHemjee311Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Малд авсан дэд арга хэмжээ",
		Identity: "id",
		Table:    "lut_avsan_arga_ded_hemjee",
		Model:    new(formModels.LutAvsanArgaDedHemjee311),
		FieldTypes: map[string]string{
			"id":                    "Text",
			"avsan_arga_hemjee_id":  "Select",
			"avsan_arga_ded_hemjee": "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"avsan_arga_hemjee_id":  []string{},
			"avsan_arga_ded_hemjee": []string{}},
		ValidationMessages: govalidator.MapData{

			"avsan_arga_hemjee_id":  []string{},
			"avsan_arga_ded_hemjee": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
