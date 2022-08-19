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

func GisLayers276Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "Давхарга",
		Identity: "id",
		Table:    "gis_layers",
		Model:    new(formModels.GisLayers276),
		FieldTypes: map[string]string{
			"id":            "Text",
			"category_id":   "Select",
			"active":        "Checkbox",
			"show":          "Checkbox",
			"name":          "Text",
			"menu_order":    "Number",
			"layer_order":   "Number",
			"layer_type":    "Radio",
			"layer_url":     "arcGISLayer",
			"info_template": "CK",
			"search_info":   "Textarea",
			"created_at":    "Text",
			"updated_at":    "Text",
			"popup_fields":  "Select",
			"search_fields": "Select",
			"style_field":   "Select",
			"check_inluded": "Checkbox",
			"export":        "Checkbox",
			"gis_legends":   "SubForm",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"category_id":   []string{},
			"name":          []string{},
			"menu_order":    []string{},
			"layer_order":   []string{},
			"layer_url":     []string{},
			"info_template": []string{},
			"search_info":   []string{},
			"popup_fields":  []string{},
			"search_fields": []string{}},
		ValidationMessages: govalidator.MapData{

			"category_id":   []string{},
			"name":          []string{},
			"menu_order":    []string{},
			"layer_order":   []string{},
			"layer_url":     []string{},
			"info_template": []string{},
			"search_info":   []string{},
			"popup_fields":  []string{},
			"search_fields": []string{}},
		SubForms: []map[string]interface{}{
			map[string]interface{}{
				"connection_field": "layer_id",
				"tableTypeColumn":  "",
				"tableTypeValue":   "",
				"table":            "gis_legends",
				"parentIdentity":   "id",
				"subIdentity":      "id",
				"subForm":          GisLegends278Dataform(),
				"subFormArray":     new([]formModels.GisLegends278),
			}},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
