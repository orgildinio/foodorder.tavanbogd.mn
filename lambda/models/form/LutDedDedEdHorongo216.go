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

func LutDedDedEdHorongo216Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Эд хөрөнгийн дэд дэд төрөл",
		Identity: "id",
		Table:    "lut_ded_ded_ed_horongo",
		Model:    new(formModels.LutDedDedEdHorongo216),
		FieldTypes: map[string]string{
			"id":                "Text",
			"ded_ed_horongo_id": "Select",
			"ded_ded_horongo":   "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"ded_ed_horongo_id": []string{},
			"ded_ded_horongo":   []string{}},
		ValidationMessages: govalidator.MapData{

			"ded_ed_horongo_id": []string{},
			"ded_ded_horongo":   []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
