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

func SubNariinUnelgeeBagBureldehuun286Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "Нарийвчилсан үнэлгээ /Багийн бүрэлдэхүүн FORM /",
		Identity: "id",
		Table:    "sub_nariin_unelgee_bag_bureldehuun",
		Model:    new(formModels.SubNariinUnelgeeBagBureldehuun286),
		FieldTypes: map[string]string{
			"id":                    "Text",
			"nariin_unelgee_id":     "",
			"ovog":                  "Text",
			"ner":                   "Text",
			"baiguullaga":           "Text",
			"register":              "Text",
			"bolovsrol":             "Text",
			"mergejil":              "Text",
			"turshilga":             "CK",
			"mail":                  "Text",
			"utas":                  "Text",
			"alban_tushaal":         "Text",
			"gamshig_sudalgaa_eseh": "Checkbox",
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
