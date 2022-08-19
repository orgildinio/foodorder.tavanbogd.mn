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

func Banner60Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "Баннер",
		Identity: "id",
		Table:    "banner",
		Model:    new(formModels.Banner60),
		FieldTypes: map[string]string{
			"id":             "Text",
			"banner_type_id": "Radio",
			"banner":         "Image",
			"news_id":        "Select",
			"created_at":     "Text",
			"updated_at":     "Text",
			"deleted_at":     "Text",
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
