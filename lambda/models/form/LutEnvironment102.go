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

func LutEnvironment102Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Алдагдсан орчин",
		Identity: "id",
		Table:    "lut_environment",
		Model:    new(formModels.LutEnvironment102),
		FieldTypes: map[string]string{
			"id":          "Text",
			"environment": "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"environment": []string{}},
		ValidationMessages: govalidator.MapData{

			"environment": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
