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

func HolbooBarih58Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "Холбоо барих",
		Identity: "id",
		Table:    "holboo_barih",
		Model:    new(formModels.HolbooBarih58),
		FieldTypes: map[string]string{
			"id":         "Text",
			"b_ner":      "Text",
			"hayag":      "Textarea",
			"bairshil":   "Geographic",
			"email":      "Text",
			"utas":       "Text",
			"created_at": "Text",
			"updated_at": "Text",
			"deleted_at": "Text",
			"web":        "Text",
			"sub_social": "SubForm",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"b_ner":    []string{},
			"hayag":    []string{},
			"bairshil": []string{},
			"email":    []string{},
			"utas":     []string{}},
		ValidationMessages: govalidator.MapData{

			"b_ner":    []string{},
			"hayag":    []string{},
			"bairshil": []string{},
			"email":    []string{},
			"utas":     []string{}},
		SubForms: []map[string]interface{}{
			map[string]interface{}{
				"connection_field": "holboo_barih_id",
				"tableTypeColumn":  "",
				"tableTypeValue":   "",
				"table":            "sub_social",
				"parentIdentity":   "id",
				"subIdentity":      "id",
				"subForm":          SubSocialHolbooBarih58Dataform(),
				"subFormArray":     new([]formModels.SubSocialHolbooBarih58),
			}},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}

func SubSocialHolbooBarih58Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "Олон нийтийн суваг",
		Identity: "id",
		Table:    "sub_social",
		Model:    new(formModels.SubSocialHolbooBarih58),
	}
}
