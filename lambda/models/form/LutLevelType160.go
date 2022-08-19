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

func LutLevelType160Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Эрсдлийн түвшин",
		Identity: "id",
		Table:    "lut_level_type",
		Model:    new(formModels.LutLevelType160),
		FieldTypes: map[string]string{
			"id":         "Text",
			"level_type": "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"level_type": []string{}},
		ValidationMessages: govalidator.MapData{

			"level_type": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
