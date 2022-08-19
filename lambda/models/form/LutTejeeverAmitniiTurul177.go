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

func LutTejeeverAmitniiTurul177Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Тэжээвэр амьтны төрөл",
		Identity: "id",
		Table:    "lut_tejeever_amitnii_turul",
		Model:    new(formModels.LutTejeeverAmitniiTurul177),
		FieldTypes: map[string]string{
			"id":                     "Text",
			"tejeever_amitnii_turul": "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"tejeever_amitnii_turul": []string{}},
		ValidationMessages: govalidator.MapData{

			"tejeever_amitnii_turul": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
