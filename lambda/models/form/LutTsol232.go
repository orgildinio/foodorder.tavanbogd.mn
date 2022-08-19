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

func LutTsol232Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "Цол",
		Identity: "id",
		Table:    "lut_tsol",
		Model:    new(formModels.LutTsol232),
		FieldTypes: map[string]string{
			"id":   "Text",
			"tsol": "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"tsol": []string{}},
		ValidationMessages: govalidator.MapData{

			"tsol": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
