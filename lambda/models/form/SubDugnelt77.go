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

func SubDugnelt77Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "form Дүгнэлт",
		Identity: "id",
		Table:    "sub_dugnelt",
		Model:    new(formModels.SubDugnelt77),
		FieldTypes: map[string]string{
			"id":              "Text",
			"danger_type_id":  "Select",
			"zowlomj":         "Textarea",
			"z_heregjilt":     "Textarea",
			"unelgee_form_id": "",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"danger_type_id": []string{},
			"zowlomj":        []string{},
			"z_heregjilt":    []string{}},
		ValidationMessages: govalidator.MapData{

			"danger_type_id": []string{},
			"zowlomj":        []string{},
			"z_heregjilt":    []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
