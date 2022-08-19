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

func LutMAccidentType321Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Гамшиг, аюулт үзэгдэл, осолд өртсөн байдал /M/",
		Identity: "id",
		Table:    "lut_m_accident_type",
		Model:    new(formModels.LutMAccidentType321),
		FieldTypes: map[string]string{
			"id":              "Text",
			"m_accident_type": "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"m_accident_type": []string{}},
		ValidationMessages: govalidator.MapData{

			"m_accident_type": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
