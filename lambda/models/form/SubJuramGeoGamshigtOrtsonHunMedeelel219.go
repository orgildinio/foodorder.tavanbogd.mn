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

func SubJuramGeoGamshigtOrtsonHunMedeelel219Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "Журам-Биологийн  form Гамшиг, аюулт үзэгдэл, осолд нэрвэгдсэн хүний мэдээлэл",
		Identity: "id",
		Table:    "sub_juram_geo_gamshigt_ortson_hun_medeelel",
		Model:    new(formModels.SubJuramGeoGamshigtOrtsonHunMedeelel219),
		FieldTypes: map[string]string{
			"id":                                "Text",
			"juram_geo_id":                      "Text",
			"ovog":                              "Text",
			"ner":                               "Text",
			"register":                          "Register",
			"nas":                               "Text",
			"huis":                              "Text",
			"jiremsen_eseh":                     "Checkbox",
			"hb_eseh":                           "Checkbox",
			"hogjliin_berhsheel_id":             "Select",
			"niigmiin_baidal_id":                "Select",
			"gamshigt_ortson_baidal_hun_id":     "Select",
			"gamshigt_ortson_baidal_materal_id": "Select",
			"us_id":                             "Text",
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
