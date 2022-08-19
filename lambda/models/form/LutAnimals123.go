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

func LutAnimals123Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Мал амьтны төрөл",
		Identity: "id",
		Table:    "lut_animals",
		Model:    new(formModels.LutAnimals123),
		FieldTypes: map[string]string{
			"id":      "Text",
			"animals": "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"animals": []string{}},
		ValidationMessages: govalidator.MapData{

			"animals": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
