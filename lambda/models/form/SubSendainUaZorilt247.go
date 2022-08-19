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

func SubSendainUaZorilt247Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "Сендайн үйл ажиллагаа /FORM",
		Identity: "id",
		Table:    "sub_sendain_ua_zorilt",
		Model:    new(formModels.SubSendainUaZorilt247),
		FieldTypes: map[string]string{
			"id":            "Text",
			"sendain_ua_id": "",
			"zorilt_mn":     "Text",
			"zorilt_en":     "Text",
			"tailbar_mn":    "CK",
			"tailbar_en":    "CK",
		},
		Formulas:           []models.Formula{},
		ValidationRules:    govalidator.MapData{},
		ValidationMessages: govalidator.MapData{},
		SubForms:           []map[string]interface{}{},
		AfterInsert:        nil,
		AfterUpdate:        nil,
		BeforeInsert:       nil,
		BeforeUpdate:       nil,
		TriggerNameSpace:   "",
	}
}
