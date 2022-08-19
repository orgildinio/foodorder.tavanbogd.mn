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

func SubJuramGeoHeregtseeToim151Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "/Журам-Геологийн/ FORM Хэрэгцээний тойм ",
		Identity: "id",
		Table:    "sub_juram_geo_heregtsee_toim",
		Model:    new(formModels.SubJuramGeoHeregtseeToim151),
		FieldTypes: map[string]string{
			"id":                       "Text",
			"juram_geo_id":             "Text",
			"heregtsee_torol_id":       "Select",
			"heregtsee_ded_torol_id":   "Select",
			"heregtsee_d_ded_torol_id": "Select",
			"too":                      "Number",
			"hemjih_negj_id":           "Select",
			"us_id":                    "Text",
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
