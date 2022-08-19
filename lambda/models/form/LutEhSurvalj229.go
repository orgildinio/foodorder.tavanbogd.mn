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

func LutEhSurvalj229Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Эх сурвалж АҮБ",
		Identity: "id",
		Table:    "lut_eh_survalj",
		Model:    new(formModels.LutEhSurvalj229),
		FieldTypes: map[string]string{
			"id":         "Text",
			"eh_survalj": "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"eh_survalj": []string{}},
		ValidationMessages: govalidator.MapData{

			"eh_survalj": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
