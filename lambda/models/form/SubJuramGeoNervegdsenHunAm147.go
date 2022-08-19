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

func SubJuramGeoNervegdsenHunAm147Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "/Журам-Геологийн/ FORM Нэрвэгдсээн хүн амын мэдээлэл 1",
		Identity: "id",
		Table:    "sub_juram_geo_nervegdsen_hun_am",
		Model:    new(formModels.SubJuramGeoNervegdsenHunAm147),
		FieldTypes: map[string]string{
			"id":                            "Text",
			"juram_geo_id":                  "Text",
			"nervegdesen_hun_id":            "Select",
			"nervegdsen_baiguullaga_id":     "Select",
			"nervegdsen_ded_baiguullaga_id": "Select",
			"hun_am_torol_id":               "Select",
			"hun_am_ded_torol_id":           "Select",
			"too":                           "Number",
			"hun_am_ded_ded_torol_id":       "Select",
			"us_id":                         "Text",
			"bugd":                          "Number",
			"hun_am_too":                    "Number",
			"hunamded_too":                  "Number",
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
