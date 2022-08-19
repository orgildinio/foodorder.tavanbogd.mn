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

func LutRatingType38Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Эрсдлийн зэрэглэл",
		Identity: "id",
		Table:    "lut_rating_type",
		Model:    new(formModels.LutRatingType38),
		FieldTypes: map[string]string{
			"id":                 "Text",
			"ersdeliin_zereglel": "Text",
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
