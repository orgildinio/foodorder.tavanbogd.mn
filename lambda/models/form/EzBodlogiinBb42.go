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

func EzBodlogiinBb42Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "Эрх зүй бодлогын баримт бичиг",
		Identity: "id",
		Table:    "ez_bodlogiin_bb",
		Model:    new(formModels.EzBodlogiinBb42),
		FieldTypes: map[string]string{
			"id":                     "Text",
			"document_sorting_id":    "Select",
			"sub_document_sorting":   "Select",
			"ez_barimt_bichgiin_ner": "Text",
			"hawsralt":               "File",
			"created_at":             "Text",
			"updated_at":             "Text",
			"deleted_at":             "Text",
			"barimt_ner_en":          "Text",
			"file_en":                "File",
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
