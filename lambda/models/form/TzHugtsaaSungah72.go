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

func TzHugtsaaSungah72Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "Тусгай зөвшөөрлийн хугацаа сунгах",
		Identity: "id",
		Table:    "tz_hugtsaa_sungah",
		Model:    new(formModels.TzHugtsaaSungah72),
		FieldTypes: map[string]string{
			"id":                         "Text",
			"b_id":                       "Select",
			"sungah_eseh":                "Radio",
			"sungah_shaltgaan":           "CK",
			"tz_gerchilgee_dugaar":       "Text",
			"tz_gerchilgee":              "File",
			"confirm_status_id":          "Select",
			"confirm_status_id_by_admin": "Select",
			"shiidwer":                   "Text",
			"shiidwer_hawsralt":          "File",
			"sungasan_ognoo":             "Date",
			"huchintei_ognoo":            "Date",
			"tailbar":                    "CK",
			"created_at":                 "Text",
			"updated_at":                 "Text",
			"deleted_at":                 "Text",
			"sub_attach_file":            "SubForm",
		},
		Formulas:           []models.Formula{},
		ValidationRules:    govalidator.MapData{},
		ValidationMessages: govalidator.MapData{},
		SubForms: []map[string]interface{}{
			map[string]interface{}{
				"connection_field": "tz_sungah_id",
				"tableTypeColumn":  "",
				"tableTypeValue":   "",
				"table":            "sub_attach_file",
				"parentIdentity":   "id",
				"subIdentity":      "id",
				"subForm":          SubAttachFileTzHugtsaaSungah72Dataform(),
				"subFormArray":     new([]formModels.SubAttachFileTzHugtsaaSungah72),
			}},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}

func SubAttachFileTzHugtsaaSungah72Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "Хавсралт",
		Identity: "id",
		Table:    "sub_attach_file",
		Model:    new(formModels.SubAttachFileTzHugtsaaSungah72),
	}
}
