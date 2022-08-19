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

func LutSubRiskreduction93Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Гамшгийн эрсдэлийг бууруулах дэд зөвлөл ",
		Identity: "id",
		Table:    "lut_sub_riskreduction",
		Model:    new(formModels.LutSubRiskreduction93),
		FieldTypes: map[string]string{
			"id":                "Text",
			"riskreduction_id":  "Select",
			"sub_riskreduction": "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"riskreduction_id":  []string{},
			"sub_riskreduction": []string{}},
		ValidationMessages: govalidator.MapData{

			"riskreduction_id":  []string{},
			"sub_riskreduction": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
