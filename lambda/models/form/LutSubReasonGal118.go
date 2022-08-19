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

func LutSubReasonGal118Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Галын Үзүүлэлт ",
		Identity: "id",
		Table:    "lut_sub_reason_gal",
		Model:    new(formModels.LutSubReasonGal118),
		FieldTypes: map[string]string{
			"id":               "Text",
			"shaltgaan_gal_id": "Select",
			"sub_reason_gal":   "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"shaltgaan_gal_id": []string{},
			"sub_reason_gal":   []string{}},
		ValidationMessages: govalidator.MapData{

			"shaltgaan_gal_id": []string{},
			"sub_reason_gal":   []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
