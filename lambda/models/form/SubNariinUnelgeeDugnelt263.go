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

func SubNariinUnelgeeDugnelt263Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "Нарийвчилсан үнэлгээ FORM",
		Identity: "id",
		Table:    "sub_nariin_unelgee_dugnelt",
		Model:    new(formModels.SubNariinUnelgeeDugnelt263),
		FieldTypes: map[string]string{
			"id":                "Text",
			"nariin_unelgee_id": "Text",
			"ayul_torol_id":     "Select",
			"zovlomj":           "CK",
			"zovlomj_heregjilt": "Number",
		},
		Formulas:           []models.Formula{},
		ValidationRules:    govalidator.MapData{},
		ValidationMessages: govalidator.MapData{},
		SubForms:           []map[string]interface{}{},
		AfterInsert:        nil,
		AfterUpdate:        nil,
		BeforeInsert:       nil,
		BeforeUpdate:       nil,
		TriggerNameSpace:   "",
	}
}
