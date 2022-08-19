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

func LutSubSubNervgdsenHun144Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Нэрвэгдсэн хүн амын дэд дэд мэдээлэл",
		Identity: "id",
		Table:    "lut_sub_sub_nervgdsen_hun",
		Model:    new(formModels.LutSubSubNervgdsenHun144),
		FieldTypes: map[string]string{
			"id":             "Text",
			"n_ded_torol_id": "Select",
			"ded_ded_torol":  "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"n_ded_torol_id": []string{},
			"ded_ded_torol":  []string{}},
		ValidationMessages: govalidator.MapData{

			"n_ded_torol_id": []string{},
			"ded_ded_torol":  []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
