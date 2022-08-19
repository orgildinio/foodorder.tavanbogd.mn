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

func LutAngi230Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "Анги",
		Identity: "id",
		Table:    "lut_angi",
		Model:    new(formModels.LutAngi230),
		FieldTypes: map[string]string{
			"id":   "Text",
			"angi": "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"angi": []string{}},
		ValidationMessages: govalidator.MapData{

			"angi": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
