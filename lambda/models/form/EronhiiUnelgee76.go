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

func EronhiiUnelgee76Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "Ерөнхий үнэлгээний тайлан",
		Identity: "id",
		Table:    "eronhii_unelgee",
		Model:    new(formModels.EronhiiUnelgee76),
		FieldTypes: map[string]string{
			"id":                        "Text",
			"user_id":                   "Select",
			"general_evalution_type_id": "Select",
			"evalution_sector_id":       "Select",
			"aimag_id":                  "Select",
			"sum_id":                    "Select",
			"bag_id":                    "Select",
			"bairshil":                  "Geographic",
			"danger_type_id":            "Select",
			"ognoo":                     "Date",
			"indicator_type_id":         "Select",
			"indicator_id":              "Select",
			"sub_indicator_id":          "Select",
			"zeregleliin_hemjee":        "Text",
			"owog":                      "Text",
			"ner":                       "Text",
			"baiguullaga":               "Text",
			"a_tushaal":                 "Text",
			"ur_dun":                    "CK",
			"zurag":                     "Image",
			"u_hiij_ehelsen_ognoo":      "Date",
			"u_hiij_duussan_ognoo":      "Date",
			"confirm_status_id":         "Select",
			"created_at":                "Text",
			"updated_at":                "Text",
			"deleted_at":                "Text",
			"sub_ersdel_unelgee":        "SubForm",
			"sub_ersdel_index":          "SubForm",
			"sub_dugnelt":               "SubForm",
			"sub_hawsral_bb":            "SubForm",
			"sub_yronhiiunelgee_bag":    "SubForm",
		},
		Formulas:           []models.Formula{},
		ValidationRules:    govalidator.MapData{},
		ValidationMessages: govalidator.MapData{},
		SubForms: []map[string]interface{}{
			map[string]interface{}{
				"connection_field": "unelgee_form_id",
				"tableTypeColumn":  "",
				"tableTypeValue":   "",
				"table":            "sub_ersdel_unelgee",
				"parentIdentity":   "id",
				"subIdentity":      "id",
				"subForm":          SubErsdelUnelgeeEronhiiUnelgee76Dataform(),
				"subFormArray":     new([]formModels.SubErsdelUnelgeeEronhiiUnelgee76),
			},
			map[string]interface{}{
				"connection_field": "unelgee_form_id",
				"tableTypeColumn":  "",
				"tableTypeValue":   "",
				"table":            "sub_ersdel_index",
				"parentIdentity":   "id",
				"subIdentity":      "id",
				"subForm":          SubErsdelIndexEronhiiUnelgee76Dataform(),
				"subFormArray":     new([]formModels.SubErsdelIndexEronhiiUnelgee76),
			},
			map[string]interface{}{
				"connection_field": "unelgee_form_id",
				"tableTypeColumn":  "",
				"tableTypeValue":   "",
				"table":            "sub_dugnelt",
				"parentIdentity":   "id",
				"subIdentity":      "id",
				"subForm":          SubDugnelt77Dataform(),
				"subFormArray":     new([]formModels.SubDugnelt77),
			},
			map[string]interface{}{
				"connection_field": "unelgee_form_id",
				"tableTypeColumn":  "",
				"tableTypeValue":   "",
				"table":            "sub_hawsral_bb",
				"parentIdentity":   "id",
				"subIdentity":      "id",
				"subForm":          SubHawsralBbEronhiiUnelgee76Dataform(),
				"subFormArray":     new([]formModels.SubHawsralBbEronhiiUnelgee76),
			},
			map[string]interface{}{
				"connection_field": "eronhii_unelgeeid",
				"tableTypeColumn":  "",
				"tableTypeValue":   "",
				"table":            "sub_yronhiiunelgee_bag",
				"parentIdentity":   "id",
				"subIdentity":      "id",
				"subForm":          SubYronhiiunelgeeBagEronhiiUnelgee76Dataform(),
				"subFormArray":     new([]formModels.SubYronhiiunelgeeBagEronhiiUnelgee76),
			}},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}

func SubErsdelUnelgeeEronhiiUnelgee76Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "Үнэлгээний үзүүлэлт",
		Identity: "id",
		Table:    "sub_ersdel_unelgee",
		Model:    new(formModels.SubErsdelUnelgeeEronhiiUnelgee76),
	}
}

func SubErsdelIndexEronhiiUnelgee76Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "Эрсдлийн индекс",
		Identity: "id",
		Table:    "sub_ersdel_index",
		Model:    new(formModels.SubErsdelIndexEronhiiUnelgee76),
	}
}

func SubHawsralBbEronhiiUnelgee76Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "Хавсралт файл",
		Identity: "id",
		Table:    "sub_hawsral_bb",
		Model:    new(formModels.SubHawsralBbEronhiiUnelgee76),
	}
}

func SubYronhiiunelgeeBagEronhiiUnelgee76Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "",
		Identity: "id",
		Table:    "sub_yronhiiunelgee_bag",
		Model:    new(formModels.SubYronhiiunelgeeBagEronhiiUnelgee76),
	}
}
