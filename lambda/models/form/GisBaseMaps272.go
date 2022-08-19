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

func GisBaseMaps272Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "Суурь зураг",
		Identity: "id",
		Table:    "gis_base_maps",
		Model:    new(formModels.GisBaseMaps272),
		FieldTypes: map[string]string{
			"id":         "Text",
			"layerName":  "Text",
			"url":        "Text",
			"minZoom":    "Number",
			"maxZoom":    "Text",
			"show":       "Checkbox",
			"is_dynamic": "Checkbox",
			"image":      "Image",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"layerName":  []string{},
			"url":        []string{},
			"minZoom":    []string{},
			"maxZoom":    []string{},
			"show":       []string{},
			"is_dynamic": []string{},
			"image":      []string{}},
		ValidationMessages: govalidator.MapData{

			"layerName":  []string{},
			"url":        []string{},
			"minZoom":    []string{},
			"maxZoom":    []string{},
			"show":       []string{},
			"is_dynamic": []string{},
			"image":      []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
