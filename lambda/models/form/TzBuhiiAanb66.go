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

func TzBuhiiAanb66Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "ТЗ БҮХИЙ ААНБ",
		Identity: "id",
		Table:    "tz_buhii_aanb",
		Model:    new(formModels.TzBuhiiAanb66),
		FieldTypes: map[string]string{
			"id":                        "Text",
			"b_ner":                     "Text",
			"ub_gerchilgee_dug":         "Text",
			"register":                  "Text",
			"ub_gerchilgee":             "File",
			"activity_type_id":          "Select",
			"evalution_sector_id":       "Select",
			"zahiral":                   "Text",
			"zaki_utas":                 "Text",
			"utas":                      "Text",
			"email":                     "Text",
			"aimag_id":                  "Select",
			"sum_id":                    "Select",
			"bag_id":                    "Select",
			"gudamj":                    "Textarea",
			"bair":                      "Textarea",
			"hayag_delegrengui":         "CK",
			"tz_dugaar":                 "Text",
			"tz_olgoson":                "Date",
			"tz_duusah":                 "Date",
			"tz_shiidwer":               "Text",
			"shiidwer_hawsralt":         "File",
			"created_at":                "Text",
			"updated_at":                "Text",
			"deleted_at":                "Text",
			"sub_tz_aan_unelgee_salbar": "SubForm",
		},
		Formulas:           []models.Formula{},
		ValidationRules:    govalidator.MapData{},
		ValidationMessages: govalidator.MapData{},
		SubForms: []map[string]interface{}{
			map[string]interface{}{
				"connection_field": "aan_id",
				"tableTypeColumn":  "",
				"tableTypeValue":   "",
				"table":            "sub_tz_aan_unelgee_salbar",
				"parentIdentity":   "id",
				"subIdentity":      "id",
				"subForm":          SubTzAanUnelgeeSalbarTzBuhiiAanb66Dataform(),
				"subFormArray":     new([]formModels.SubTzAanUnelgeeSalbarTzBuhiiAanb66),
			}},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}

func SubTzAanUnelgeeSalbarTzBuhiiAanb66Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "",
		Identity: "id",
		Table:    "sub_tz_aan_unelgee_salbar",
		Model:    new(formModels.SubTzAanUnelgeeSalbarTzBuhiiAanb66),
	}
}
