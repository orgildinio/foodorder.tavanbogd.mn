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

func SubJuramGeoAvrahSergeeh149Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "/Журам-Геологийн/ FORM Аврах, хор уршгийг арилгах, хойшлуулшгүй сэргээн босгох үйл ажиллагаанд ажилласан хүч хэрэгсэл, зарцуулсан зардал",
		Identity: "id",
		Table:    "sub_juram_geo_avrah_sergeeh",
		Model:    new(formModels.SubJuramGeoAvrahSergeeh149),
		FieldTypes: map[string]string{
			"id":                    "Text",
			"juram_geo_id":          "Text",
			"undsen_torol_id":       "Select",
			"undsen_ded_torol_id":   "Select",
			"baiguullaga_torol_id":  "Select",
			"too":                   "Number",
			"mongo":                 "Number",
			"niit_too":              "Number",
			"sub_baiguullaga_torol": "SubForm",
			"us_id":                 "Text",
		},
		Formulas:           []models.Formula{},
		ValidationRules:    govalidator.MapData{},
		ValidationMessages: govalidator.MapData{},
		SubForms: []map[string]interface{}{
			map[string]interface{}{
				"connection_field": "juram_geo_avrah_sergeeh_id",
				"tableTypeColumn":  "",
				"tableTypeValue":   "",
				"table":            "sub_baiguullaga_torol",
				"parentIdentity":   "id",
				"subIdentity":      "id",
				"subForm":          SubBaiguullagaTorolSubJuramGeoAvrahSergeeh149Dataform(),
				"subFormArray":     new([]formModels.SubBaiguullagaTorolSubJuramGeoAvrahSergeeh149),
			}},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}

func SubBaiguullagaTorolSubJuramGeoAvrahSergeeh149Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "Байгууллагын төрөл ",
		Identity: "id",
		Table:    "sub_baiguullaga_torol",
		Model:    new(formModels.SubBaiguullagaTorolSubJuramGeoAvrahSergeeh149),
	}
}
