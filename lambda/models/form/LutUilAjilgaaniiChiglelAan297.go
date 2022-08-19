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

func LutUilAjilgaaniiChiglelAan297Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L ТЗ БҮХИЙ ААН үйл ажилгааны төрөл",
		Identity: "id",
		Table:    "lut_uil_ajilgaanii_chiglel_aan",
		Model:    new(formModels.LutUilAjilgaaniiChiglelAan297),
		FieldTypes: map[string]string{
			"id":                     "Text",
			"uil_ajilgaanii_chiglel": "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"uil_ajilgaanii_chiglel": []string{}},
		ValidationMessages: govalidator.MapData{

			"uil_ajilgaanii_chiglel": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
