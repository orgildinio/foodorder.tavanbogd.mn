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

func DundHugtsaaStrategy54Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "Дунд хугацааны стратеги",
		Identity: "id",
		Table:    "dund_hugtsaa_strategy",
		Model:    new(formModels.DundHugtsaaStrategy54),
		FieldTypes: map[string]string{
			"id":                        "Text",
			"created_at":                "Text",
			"updated_at":                "Text",
			"deleted_at":                "Text",
			"mongol_strategy":           "File",
			"sub_dund_hugtsaa_strategy": "SubForm",
		},
		Formulas:           []models.Formula{},
		ValidationRules:    govalidator.MapData{},
		ValidationMessages: govalidator.MapData{},
		SubForms: []map[string]interface{}{
			map[string]interface{}{
				"connection_field": "dund_hugtsaa_strategy_id",
				"tableTypeColumn":  "",
				"tableTypeValue":   "",
				"table":            "sub_dund_hugtsaa_strategy",
				"parentIdentity":   "id",
				"subIdentity":      "id",
				"subForm":          SubDundHugtsaaStrategyDundHugtsaaStrategy54Dataform(),
				"subFormArray":     new([]formModels.SubDundHugtsaaStrategyDundHugtsaaStrategy54),
			}},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}

func SubDundHugtsaaStrategyDundHugtsaaStrategy54Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "",
		Identity: "id",
		Table:    "sub_dund_hugtsaa_strategy",
		Model:    new(formModels.SubDundHugtsaaStrategyDundHugtsaaStrategy54),
	}
}
