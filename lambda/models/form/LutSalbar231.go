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

func LutSalbar231Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "Салбар",
		Identity: "id",
		Table:    "lut_salbar",
		Model:    new(formModels.LutSalbar231),
		FieldTypes: map[string]string{
			"id":      "Text",
			"angi_id": "Select",
			"salbar":  "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"angi_id": []string{},
			"salbar":  []string{}},
		ValidationMessages: govalidator.MapData{

			"angi_id": []string{},
			"salbar":  []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
