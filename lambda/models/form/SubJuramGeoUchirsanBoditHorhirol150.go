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

func SubJuramGeoUchirsanBoditHorhirol150Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "/Журам-Геологийн/ FORM Учирсан бодит хохирол ",
		Identity: "id",
		Table:    "sub_juram_geo_uchirsan_bodit_horhirol",
		Model:    new(formModels.SubJuramGeoUchirsanBoditHorhirol150),
		FieldTypes: map[string]string{
			"id":                     "Text",
			"juram_geo_id":           "Text",
			"horogdson_baidal_id":    "Select",
			"hohirol_torol_id":       "Select",
			"hohirol_ded_torol_id":   "Select",
			"hohirol_d_ded_torol_id": "Select",
			"hohirol_baidal_id":      "Select",
			"too":                    "Number",
			"hemjih_negj_id":         "Select",
			"uchirsan_hohirol":       "Number",
			"us_id":                  "Text",
			"hohirol_too":            "Number",
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
