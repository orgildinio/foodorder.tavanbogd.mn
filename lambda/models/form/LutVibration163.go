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

func LutVibration163Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Чичирхийллийн хэмжээ",
		Identity: "id",
		Table:    "lut_vibration",
		Model:    new(formModels.LutVibration163),
		FieldTypes: map[string]string{
			"id":        "Text",
			"vibration": "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"vibration": []string{}},
		ValidationMessages: govalidator.MapData{

			"vibration": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
