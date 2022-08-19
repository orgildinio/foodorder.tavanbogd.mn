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

func LutAlbanTushaal165Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "Албан тушаал",
		Identity: "id",
		Table:    "lut_alban_tushaal",
		Model:    new(formModels.LutAlbanTushaal165),
		FieldTypes: map[string]string{
			"id":            "Text",
			"alban_tushaal": "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"alban_tushaal": []string{}},
		ValidationMessages: govalidator.MapData{

			"alban_tushaal": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
