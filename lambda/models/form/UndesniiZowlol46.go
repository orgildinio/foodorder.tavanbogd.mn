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

func UndesniiZowlol46Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "Үндэсний зөвлөл",
		Identity: "id",
		Table:    "undesnii_zowlol",
		Model:    new(formModels.UndesniiZowlol46),
		FieldTypes: map[string]string{
			"id":                         "Text",
			"butets_bureldhuun":          "Image",
			"created_at":                 "Text",
			"updated_at":                 "Text",
			"deleted_at":                 "Text",
			"sub_undes_zovlol_huraldaan": "SubForm",
			"sub_undes_zovlol_gamshig_ersdel_buuruulah": "SubForm",
		},
		Formulas:           []models.Formula{},
		ValidationRules:    govalidator.MapData{},
		ValidationMessages: govalidator.MapData{},
		SubForms: []map[string]interface{}{
			map[string]interface{}{
				"connection_field": "undes_zovlol_id",
				"tableTypeColumn":  "",
				"tableTypeValue":   "",
				"table":            "sub_undes_zovlol_huraldaan",
				"parentIdentity":   "id",
				"subIdentity":      "id",
				"subForm":          SubUndesZovlolHuraldaanUndesniiZowlol46Dataform(),
				"subFormArray":     new([]formModels.SubUndesZovlolHuraldaanUndesniiZowlol46),
			},
			map[string]interface{}{
				"connection_field": "undes_zovlol_id",
				"tableTypeColumn":  "",
				"tableTypeValue":   "",
				"table":            "sub_undes_zovlol_gamshig_ersdel_buuruulah",
				"parentIdentity":   "id",
				"subIdentity":      "id",
				"subForm":          SubUndesZovlolGamshigErsdelBuuruulahUndesniiZowlol46Dataform(),
				"subFormArray":     new([]formModels.SubUndesZovlolGamshigErsdelBuuruulahUndesniiZowlol46),
			}},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}

func SubUndesZovlolHuraldaanUndesniiZowlol46Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "Хуралдаан",
		Identity: "id",
		Table:    "sub_undes_zovlol_huraldaan",
		Model:    new(formModels.SubUndesZovlolHuraldaanUndesniiZowlol46),
	}
}

func SubUndesZovlolGamshigErsdelBuuruulahUndesniiZowlol46Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "Бүрэлдэхүүн",
		Identity: "id",
		Table:    "sub_undes_zovlol_gamshig_ersdel_buuruulah",
		Model:    new(formModels.SubUndesZovlolGamshigErsdelBuuruulahUndesniiZowlol46),
	}
}
