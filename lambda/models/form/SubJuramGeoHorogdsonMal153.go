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

func SubJuramGeoHorogdsonMal153Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "/Журам-Геологийн/ FORM Хорогдсон мал нас хүйсээр ",
		Identity: "id",
		Table:    "sub_juram_geo_horogdson_mal",
		Model:    new(formModels.SubJuramGeoHorogdsonMal153),
		FieldTypes: map[string]string{
			"id":                         "Text",
			"juram_goe_id":               "Text",
			"mal_torol_id":               "Select",
			"mal_ded_torol_id":           "Select",
			"mal_d_ded_torol_id":         "Select",
			"huis":                       "Select",
			"avsan_arga_hemjee_id":       "Select",
			"niit_ovchilson_mal":         "Number",
			"niit_horogdson_mal":         "Number",
			"niit_ovchilson_mal_torloor": "Select",
			"niit_horogdson_mal_torloor": "Select",
			"us_id":                      "Text",
			"mal_too":                    "Number",
			"huis_too":                   "Number",
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
