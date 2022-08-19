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

func HongiinMedee325Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     " Хоногийн мэдээ ",
		Identity: "id",
		Table:    "hongiin_medee",
		Model:    new(formModels.HongiinMedee325),
		FieldTypes: map[string]string{
			"id":                    "Text",
			"ayul_torol_id":         "Select",
			"ayul_ded_torol_id":     "Select",
			"ayul_ded_ded_torol_id": "Select",
			"aimagid":               "Select",
			"sum_id":                "Select",
			"bairshil":              "Geographic",
			"urtrag":                "Text",
			"orgorog":               "Text",
			"m_avsan_ognoo_obeg":    "DateTime",
			"m_avsan_ognoo_salbar":  "DateTime",
			"e_ognoo":               "DateTime",
			"d_ognoo":               "DateTime",
			"ajilah_huch":           "Textarea",
			"vaktsin":               "Textarea",
			"niit_hohirol":          "Textarea",
			"avsan_arga_hemjee":     "Textarea",
			"created_at":            "Text",
			"updated_at":            "Text",
			"user_id":               "Text",
			"medeelel":              "Textarea",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"ayul_torol_id":         []string{},
			"ayul_ded_torol_id":     []string{},
			"ayul_ded_ded_torol_id": []string{},
			"aimagid":               []string{},
			"sum_id":                []string{}},
		ValidationMessages: govalidator.MapData{

			"ayul_torol_id":         []string{},
			"ayul_ded_torol_id":     []string{},
			"ayul_ded_ded_torol_id": []string{},
			"aimagid":               []string{},
			"sum_id":                []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
