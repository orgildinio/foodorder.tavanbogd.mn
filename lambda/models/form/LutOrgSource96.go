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

func LutOrgSource96Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Эх сурвалж /БГ/",
		Identity: "id",
		Table:    "lut_org_source",
		Model:    new(formModels.LutOrgSource96),
		FieldTypes: map[string]string{
			"id":         "Text",
			"org_source": "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"org_source": []string{}},
		ValidationMessages: govalidator.MapData{

			"org_source": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
