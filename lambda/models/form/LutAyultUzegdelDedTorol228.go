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

func LutAyultUzegdelDedTorol228Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Аюулт үзэгдийн дэд төрөл",
		Identity: "id",
		Table:    "lut_ayult_uzegdel_ded_torol",
		Model:    new(formModels.LutAyultUzegdelDedTorol228),
		FieldTypes: map[string]string{
			"id":                      "Text",
			"ayult_uzegdel_torol_id":  "Select",
			"ayult_uzegdel_ded_torol": "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"ayult_uzegdel_torol_id":  []string{},
			"ayult_uzegdel_ded_torol": []string{}},
		ValidationMessages: govalidator.MapData{

			"ayult_uzegdel_torol_id":  []string{},
			"ayult_uzegdel_ded_torol": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
