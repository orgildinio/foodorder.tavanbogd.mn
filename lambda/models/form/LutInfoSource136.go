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

func LutInfoSource136Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Мэдээллийг хүлээн авсан суваг",
		Identity: "id",
		Table:    "lut_info_source",
		Model:    new(formModels.LutInfoSource136),
		FieldTypes: map[string]string{
			"id":          "Text",
			"info_source": "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"info_source": []string{}},
		ValidationMessages: govalidator.MapData{

			"info_source": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
