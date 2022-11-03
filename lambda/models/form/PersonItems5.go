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

func PersonItems5Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "Эд зүйл",
		Identity: "ID",
		Table:    "PERSON_ITEMS",
		Model:    new(formModels.PersonItems5),
		FieldTypes: map[string]string{
			"ID":               "Text",
			"PERSON_ID":        "Select",
			"ITEM_NAME":        "Text",
			"ITEM_DESCRIPTION": "Textarea",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"PERSON_ID": []string{"required"}},
		ValidationMessages: govalidator.MapData{

			"PERSON_ID": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
