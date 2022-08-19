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

func News63Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "Мэдээ мэдээлэл",
		Identity: "id",
		Table:    "news",
		Model:    new(formModels.News63),
		FieldTypes: map[string]string{
			"id":                     "Text",
			"news_type_id":           "Select",
			"garchig":                "Text",
			"delgerengui":            "CK",
			"zurag":                  "Image",
			"ursah_eseh":             "Checkbox",
			"created_at":             "Text",
			"updated_at":             "Text",
			"deleted_at":             "Text",
			"handalt":                "Text",
			"logo":                   "Text",
			"share":                  "Text",
			"sub_news_social_typies": "SubForm",
		},
		Formulas:           []models.Formula{},
		ValidationRules:    govalidator.MapData{},
		ValidationMessages: govalidator.MapData{},
		SubForms: []map[string]interface{}{
			map[string]interface{}{
				"connection_field": "news_id",
				"tableTypeColumn":  "",
				"tableTypeValue":   "",
				"table":            "sub_news_social_typies",
				"parentIdentity":   "id",
				"subIdentity":      "id",
				"subForm":          SubNewsSocialTypiesNews63Dataform(),
				"subFormArray":     new([]formModels.SubNewsSocialTypiesNews63),
			}},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}

func SubNewsSocialTypiesNews63Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "Сошиал холбоос",
		Identity: "id",
		Table:    "sub_news_social_typies",
		Model:    new(formModels.SubNewsSocialTypiesNews63),
	}
}
