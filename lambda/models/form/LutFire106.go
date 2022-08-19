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

func LutFire106Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Гал түймрийг унтраахад ашигласан арга",
		Identity: "id",
		Table:    "lut_fire",
		Model:    new(formModels.LutFire106),
		FieldTypes: map[string]string{
			"id":   "Text",
			"fire": "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"fire": []string{}},
		ValidationMessages: govalidator.MapData{

			"fire": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
