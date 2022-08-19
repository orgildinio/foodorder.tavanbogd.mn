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

func GisCategory274Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "Давхарга бүлэг",
		Identity: "id",
		Table:    "gis_category",
		Model:    new(formModels.GisCategory274),
		FieldTypes: map[string]string{
			"id":          "Text",
			"name":        "Text",
			"icon":        "Image",
			"menu_order":  "Number",
			"layer_order": "Number",
			"created_at":  "Text",
			"updated_at":  "Text",
			"active":      "Checkbox",
			"show":        "Checkbox",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"name":        []string{},
			"icon":        []string{},
			"menu_order":  []string{},
			"layer_order": []string{}},
		ValidationMessages: govalidator.MapData{

			"name":        []string{},
			"icon":        []string{},
			"menu_order":  []string{},
			"layer_order": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
