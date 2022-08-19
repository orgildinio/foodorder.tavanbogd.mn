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

func SubJuramGeoEvdersenEdZuils155Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "/Журам-Геологийн/ FORM Гамшиг, аюулт үзэгдэл, ослын улмаас эвдэрч сүйдсэн эд хөрөнгийн бүртгэл",
		Identity: "id",
		Table:    "sub_juram_geo_evdersen_ed_zuils",
		Model:    new(formModels.SubJuramGeoEvdersenEdZuils155),
		FieldTypes: map[string]string{
			"id":                          "Text",
			"juram_geo_id":                "Text",
			"ed_horongo_torol_id":         "Select",
			"ed_horongo_ded_torol_id":     "Select",
			"hohirol_id":                  "Select",
			"too":                         "Number",
			"daatgal_eseh":                "Checkbox",
			"todorkhoilolt":               "Textarea",
			"ed_horongo_ded_ded_torol_id": "Select",
			"us_id":                       "Text",
			"bugd":                        "Number",
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
