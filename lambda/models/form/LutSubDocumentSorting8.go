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

func LutSubDocumentSorting8Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Эрх зүй бодлогын баримт бичиг дэд ангилал ",
		Identity: "id",
		Table:    "lut_sub_document_sorting",
		Model:    new(formModels.LutSubDocumentSorting8),
		FieldTypes: map[string]string{
			"id":                   "Text",
			"document_sorting_id":  "Select",
			"sub_document_sorting": "Text",
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
