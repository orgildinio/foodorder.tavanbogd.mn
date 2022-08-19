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

func LutNiigmiinBaidal319Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Нийгмийн байдал",
		Identity: "id",
		Table:    "lut_niigmiin_baidal",
		Model:    new(formModels.LutNiigmiinBaidal319),
		FieldTypes: map[string]string{
			"id":              "Text",
			"niigmiin_baidal": "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"niigmiin_baidal": []string{}},
		ValidationMessages: govalidator.MapData{

			"niigmiin_baidal": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
