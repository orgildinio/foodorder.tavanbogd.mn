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

func LutUndsenTorol306Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Үндсэн төрөл ",
		Identity: "id",
		Table:    "lut_undsen_torol",
		Model:    new(formModels.LutUndsenTorol306),
		FieldTypes: map[string]string{
			"id":           "Text",
			"undsen_torol": "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"undsen_torol": []string{}},
		ValidationMessages: govalidator.MapData{

			"undsen_torol": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
