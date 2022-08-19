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

func LutHorogdsonAmitan176Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Хорогдсон тэжээвэр амьтан",
		Identity: "id",
		Table:    "lut_horogdson_amitan",
		Model:    new(formModels.LutHorogdsonAmitan176),
		FieldTypes: map[string]string{
			"id":               "Text",
			"horogdson_amitan": "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"horogdson_amitan": []string{}},
		ValidationMessages: govalidator.MapData{

			"horogdson_amitan": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
