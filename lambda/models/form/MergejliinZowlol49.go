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

func MergejliinZowlol49Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "Мэргэжлийн зөвлөл",
		Identity: "id",
		Table:    "mergejliin_zowlol",
		Model:    new(formModels.MergejliinZowlol49),
		FieldTypes: map[string]string{
			"id":                              "Text",
			"huraldaan":                       "Textarea",
			"created_at":                      "Text",
			"updated_at":                      "Text",
			"deleted_at":                      "Text",
			"bureldehuun_mn":                  "Text",
			"delegation":                      "Text",
			"conference":                      "Textarea",
			"sub_mergejil_zovolgoo_tolvolgoo": "SubForm",
			"sub_mergejil_zovlol_huraldaan":   "SubForm",
			"sub_mergejil_zovlol_tushaal":     "SubForm",
		},
		Formulas:           []models.Formula{},
		ValidationRules:    govalidator.MapData{},
		ValidationMessages: govalidator.MapData{},
		SubForms: []map[string]interface{}{
			map[string]interface{}{
				"connection_field": "m_zovlol_id",
				"tableTypeColumn":  "",
				"tableTypeValue":   "",
				"table":            "sub_mergejil_zovolgoo_tolvolgoo",
				"parentIdentity":   "id",
				"subIdentity":      "id",
				"subForm":          SubMergejilZovolgooTolvolgooMergejliinZowlol49Dataform(),
				"subFormArray":     new([]formModels.SubMergejilZovolgooTolvolgooMergejliinZowlol49),
			},
			map[string]interface{}{
				"connection_field": "m_zovlol_id",
				"tableTypeColumn":  "",
				"tableTypeValue":   "",
				"table":            "sub_mergejil_zovlol_huraldaan",
				"parentIdentity":   "id",
				"subIdentity":      "id",
				"subForm":          SubMergejilZovlolHuraldaanMergejliinZowlol49Dataform(),
				"subFormArray":     new([]formModels.SubMergejilZovlolHuraldaanMergejliinZowlol49),
			},
			map[string]interface{}{
				"connection_field": "m_zovlol_id",
				"tableTypeColumn":  "",
				"tableTypeValue":   "",
				"table":            "sub_mergejil_zovlol_tushaal",
				"parentIdentity":   "id",
				"subIdentity":      "id",
				"subForm":          SubMergejilZovlolTushaalMergejliinZowlol49Dataform(),
				"subFormArray":     new([]formModels.SubMergejilZovlolTushaalMergejliinZowlol49),
			}},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}

func SubMergejilZovolgooTolvolgooMergejliinZowlol49Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "",
		Identity: "id",
		Table:    "sub_mergejil_zovolgoo_tolvolgoo",
		Model:    new(formModels.SubMergejilZovolgooTolvolgooMergejliinZowlol49),
	}
}

func SubMergejilZovlolHuraldaanMergejliinZowlol49Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "",
		Identity: "id",
		Table:    "sub_mergejil_zovlol_huraldaan",
		Model:    new(formModels.SubMergejilZovlolHuraldaanMergejliinZowlol49),
	}
}

func SubMergejilZovlolTushaalMergejliinZowlol49Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "",
		Identity: "id",
		Table:    "sub_mergejil_zovlol_tushaal",
		Model:    new(formModels.SubMergejilZovlolTushaalMergejliinZowlol49),
	}
}
