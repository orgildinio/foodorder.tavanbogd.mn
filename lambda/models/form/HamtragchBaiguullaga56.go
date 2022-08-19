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

func HamtragchBaiguullaga56Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "Хамтрагч байгууллагууд",
		Identity: "id",
		Table:    "hamtragch_baiguullaga",
		Model:    new(formModels.HamtragchBaiguullaga56),
		FieldTypes: map[string]string{
			"id":         "Text",
			"b_ner":      "Text",
			"logo":       "Image",
			"link":       "Text",
			"created_at": "Text",
			"updated_at": "Text",
			"deleted_at": "Text",
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
