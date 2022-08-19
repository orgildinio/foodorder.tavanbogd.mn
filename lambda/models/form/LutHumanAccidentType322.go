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

func LutHumanAccidentType322Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "Гамшиг, аюулт үзэгдэл, осолд өртсөн байдал /H/",
		Identity: "id",
		Table:    "lut_human_accident_type",
		Model:    new(formModels.LutHumanAccidentType322),
		FieldTypes: map[string]string{
			"id":              "Text",
			"a_accident_type": "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"a_accident_type": []string{}},
		ValidationMessages: govalidator.MapData{

			"a_accident_type": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
