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

func LutDedNervegdsenHun140Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Нэрвэгдсэн хүн амын дэд мэдээлэл ",
		Identity: "id",
		Table:    "lut_ded_nervegdsen_hun",
		Model:    new(formModels.LutDedNervegdsenHun140),
		FieldTypes: map[string]string{
			"id":                 "Text",
			"nervegdsen_hun_id":  "Select",
			"ded_nervegdsen_hun": "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"nervegdsen_hun_id":  []string{},
			"ded_nervegdsen_hun": []string{}},
		ValidationMessages: govalidator.MapData{

			"nervegdsen_hun_id":  []string{},
			"ded_nervegdsen_hun": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
