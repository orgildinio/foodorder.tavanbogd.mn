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

func SubJuramGeoTejeeverAmitan154Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "/Журам-Геологийн/ FORM Хорогдсон тэжээвэр амьтан нас хүйсээр ",
		Identity: "id",
		Table:    "sub_juram_geo_tejeever_amitan",
		Model:    new(formModels.SubJuramGeoTejeeverAmitan154),
		FieldTypes: map[string]string{
			"id":                           "Text",
			"juram_geo_id":                 "Text",
			"horogdson_tejeever_amitan_id": "Select",
			"tejeever_amitan_torol_id":     "Select",
			"too":                          "Number",
			"niit_too":                     "Number",
			"us_id":                        "Text",
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
