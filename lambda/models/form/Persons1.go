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

func Persons1Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "Хүн",
		Identity: "PERSON_ID",
		Table:    "PERSONS",
		Model:    new(formModels.Persons1),
		FieldTypes: map[string]string{
			"PERSON_ID":    "Text",
			"FIRST_NAME":   "Text",
			"LAST_NAME":    "Text",
			"CREATED_AT":   "Text",
			"UPDATED_AT":   "Text",
			"DELETED_AT":   "Text",
			"PERSON_ITEMS": "SubForm",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"FIRST_NAME": []string{"required"},
			"LAST_NAME":  []string{"required"}},
		ValidationMessages: govalidator.MapData{

			"FIRST_NAME": []string{},
			"LAST_NAME":  []string{}},
		SubForms: []map[string]interface{}{
			map[string]interface{}{
				"connection_field": "PERSON_ID",
				"tableTypeColumn":  "",
				"tableTypeValue":   "",
				"table":            "PERSON_ITEMS",
				"parentIdentity":   "PERSON_ID",
				"subIdentity":      "ID",
				"subForm":          PersonItemsPersons1Dataform(),
				"subFormArray":     new([]formModels.PersonItemsPersons1),
			}},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}

func PersonItemsPersons1Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "Эд зүйл",
		Identity: "ID",
		Table:    "PERSON_ITEMS",
		Model:    new(formModels.PersonItemsPersons1),
	}
}
