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

func SubJuramGeoNervegdsenAmitan148Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "/Журам-Геологийн/ FORM Нэрвэгдсээн хүн амын мэдээлэл 2",
		Identity: "id",
		Table:    "sub_juram_geo_nervegdsen_amitan",
		Model:    new(formModels.SubJuramGeoNervegdsenAmitan148),
		FieldTypes: map[string]string{
			"id":                            "Text",
			"juram_geo_id":                  "Text",
			"ovchilson_mal_id":              "Select",
			"ovchilson_ded_mal_id":          "Select",
			"ovchilson_amitan_torol_id":     "Select",
			"ovchilson_amitan_ded_torol_id": "Select",
			"bugd":                          "Number",
			"us_id":                         "Text",
			"mal_too":                       "Number",
			"dedmal_too":                    "Number",
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
