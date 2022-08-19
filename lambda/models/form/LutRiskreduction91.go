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

func LutRiskreduction91Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Гамшгийн эрсдэлийг бууруулах зөвлөл",
		Identity: "id",
		Table:    "lut_riskreduction",
		Model:    new(formModels.LutRiskreduction91),
		FieldTypes: map[string]string{
			"id":            "Text",
			"riskreduction": "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"riskreduction": []string{}},
		ValidationMessages: govalidator.MapData{

			"riskreduction": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
