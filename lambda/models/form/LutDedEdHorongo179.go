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

func LutDedEdHorongo179Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Эд хөрөнгийн дэд төрөл",
		Identity: "id",
		Table:    "lut_ded_ed_horongo",
		Model:    new(formModels.LutDedEdHorongo179),
		FieldTypes: map[string]string{
			"id":                  "Text",
			"ed_horongo_torol_id": "Select",
			"ded_ed_horongo":      "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"ed_horongo_torol_id": []string{},
			"ded_ed_horongo":      []string{}},
		ValidationMessages: govalidator.MapData{

			"ed_horongo_torol_id": []string{},
			"ded_ed_horongo":      []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
