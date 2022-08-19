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

func LutAjillasanBaidalGal221Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Ажилсан байдал",
		Identity: "id",
		Table:    "lut_ajillasan_baidal_gal",
		Model:    new(formModels.LutAjillasanBaidalGal221),
		FieldTypes: map[string]string{
			"id":                   "Text",
			"ajillasan baidal_gal": "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"ajillasan baidal_gal": []string{}},
		ValidationMessages: govalidator.MapData{

			"ajillasan baidal_gal": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
