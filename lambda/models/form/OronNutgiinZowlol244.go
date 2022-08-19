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

func OronNutgiinZowlol244Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "Орон нутгийн зөвлөл",
		Identity: "id",
		Table:    "oron_nutgiin_zowlol",
		Model:    new(formModels.OronNutgiinZowlol244),
		FieldTypes: map[string]string{
			"id":                            "Text",
			"created_at":                    "Text",
			"updated_at":                    "Text",
			"deleted_at":                    "Text",
			"aimag_id":                      "Select",
			"sum_id":                        "Select",
			"sub_oron_nutag_togtool":        "SubForm",
			"sub_oron_nutag_ua_tolvolgoo":   "SubForm",
			"sub_oron_nutag_ua_heregjilt":   "SubForm",
			"sub_oron_nutag_gamshig_ersdel": "SubForm",
		},
		Formulas:           []models.Formula{},
		ValidationRules:    govalidator.MapData{},
		ValidationMessages: govalidator.MapData{},
		SubForms: []map[string]interface{}{
			map[string]interface{}{
				"connection_field": "o_zovlol_id",
				"tableTypeColumn":  "",
				"tableTypeValue":   "",
				"table":            "sub_oron_nutag_togtool",
				"parentIdentity":   "id",
				"subIdentity":      "id",
				"subForm":          SubOronNutagTogtoolOronNutgiinZowlol244Dataform(),
				"subFormArray":     new([]formModels.SubOronNutagTogtoolOronNutgiinZowlol244),
			},
			map[string]interface{}{
				"connection_field": "o_zovlol_id",
				"tableTypeColumn":  "",
				"tableTypeValue":   "",
				"table":            "sub_oron_nutag_ua_tolvolgoo",
				"parentIdentity":   "id",
				"subIdentity":      "id",
				"subForm":          SubOronNutagUaTolvolgooOronNutgiinZowlol244Dataform(),
				"subFormArray":     new([]formModels.SubOronNutagUaTolvolgooOronNutgiinZowlol244),
			},
			map[string]interface{}{
				"connection_field": "o_zovlol_id",
				"tableTypeColumn":  "",
				"tableTypeValue":   "",
				"table":            "sub_oron_nutag_ua_heregjilt",
				"parentIdentity":   "id",
				"subIdentity":      "id",
				"subForm":          SubOronNutagUaHeregjiltOronNutgiinZowlol244Dataform(),
				"subFormArray":     new([]formModels.SubOronNutagUaHeregjiltOronNutgiinZowlol244),
			},
			map[string]interface{}{
				"connection_field": "o_zovlol_id",
				"tableTypeColumn":  "",
				"tableTypeValue":   "",
				"table":            "sub_oron_nutag_gamshig_ersdel",
				"parentIdentity":   "id",
				"subIdentity":      "id",
				"subForm":          SubOronNutagGamshigErsdelOronNutgiinZowlol244Dataform(),
				"subFormArray":     new([]formModels.SubOronNutagGamshigErsdelOronNutgiinZowlol244),
			}},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}

func SubOronNutagTogtoolOronNutgiinZowlol244Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "",
		Identity: "id",
		Table:    "sub_oron_nutag_togtool",
		Model:    new(formModels.SubOronNutagTogtoolOronNutgiinZowlol244),
	}
}

func SubOronNutagUaTolvolgooOronNutgiinZowlol244Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "",
		Identity: "id",
		Table:    "sub_oron_nutag_ua_tolvolgoo",
		Model:    new(formModels.SubOronNutagUaTolvolgooOronNutgiinZowlol244),
	}
}

func SubOronNutagUaHeregjiltOronNutgiinZowlol244Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "",
		Identity: "id",
		Table:    "sub_oron_nutag_ua_heregjilt",
		Model:    new(formModels.SubOronNutagUaHeregjiltOronNutgiinZowlol244),
	}
}

func SubOronNutagGamshigErsdelOronNutgiinZowlol244Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "",
		Identity: "id",
		Table:    "sub_oron_nutag_gamshig_ersdel",
		Model:    new(formModels.SubOronNutagGamshigErsdelOronNutgiinZowlol244),
	}
}
