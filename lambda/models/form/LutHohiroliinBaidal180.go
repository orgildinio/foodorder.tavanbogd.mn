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

func LutHohiroliinBaidal180Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Хохирлын байдал",
		Identity: "id",
		Table:    "lut_hohiroliin_baidal",
		Model:    new(formModels.LutHohiroliinBaidal180),
		FieldTypes: map[string]string{
			"id":                "Text",
			"hohiroliin_baidal": "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"hohiroliin_baidal": []string{}},
		ValidationMessages: govalidator.MapData{

			"hohiroliin_baidal": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
