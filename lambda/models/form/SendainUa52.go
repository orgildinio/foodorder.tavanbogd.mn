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

func SendainUa52Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "Сендайн үйл ажиллагаа",
		Identity: "id",
		Table:    "sendain_ua",
		Model:    new(formModels.SendainUa52),
		FieldTypes: map[string]string{
			"id":                    "Text",
			"ua_huree":              "File",
			"created_at":            "Text",
			"updated_at":            "Text",
			"deleted_at":            "Text",
			"sub_sendain_ua_zorilt": "SubForm",
		},
		Formulas:           []models.Formula{},
		ValidationRules:    govalidator.MapData{},
		ValidationMessages: govalidator.MapData{},
		SubForms: []map[string]interface{}{
			map[string]interface{}{
				"connection_field": "sendain_ua_id",
				"tableTypeColumn":  "",
				"tableTypeValue":   "",
				"table":            "sub_sendain_ua_zorilt",
				"parentIdentity":   "id",
				"subIdentity":      "id",
				"subForm":          SubSendainUaZorilt247Dataform(),
				"subFormArray":     new([]formModels.SubSendainUaZorilt247),
			}},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
