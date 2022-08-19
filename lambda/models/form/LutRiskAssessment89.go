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

func LutRiskAssessment89Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Мэргэжлийн зөвлөлийн бүтэц",
		Identity: "id",
		Table:    "lut_risk_assessment",
		Model:    new(formModels.LutRiskAssessment89),
		FieldTypes: map[string]string{
			"id":              "Text",
			"risk_assessment": "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"risk_assessment": []string{}},
		ValidationMessages: govalidator.MapData{

			"risk_assessment": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
