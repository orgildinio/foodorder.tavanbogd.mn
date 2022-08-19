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

func SubJuramGeoNervegdsenAmitan207Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "/Журам-Геологийн/ FORM Өвчилсөн мал амьтны мэдээлэл ",
		Identity: "id",
		Table:    "sub_juram_geo_nervegdsen_amitan",
		Model:    new(formModels.SubJuramGeoNervegdsenAmitan207),
		FieldTypes: map[string]string{
			"id":                            "Text",
			"juram_geo_id":                  "Text",
			"ovchilson_mal_id":              "Select",
			"ovchilson_ded_mal_id":          "Select",
			"ovchilson_amitan_torol_id":     "Select",
			"ovchilson_amitan_ded_torol_id": "Select",
			"bugd":                          "Text",
			"us_id":                         "Text",
			"mal_too":                       "Text",
			"dedmal_too":                    "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"bugd": []string{}},
		ValidationMessages: govalidator.MapData{

			"bugd": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
