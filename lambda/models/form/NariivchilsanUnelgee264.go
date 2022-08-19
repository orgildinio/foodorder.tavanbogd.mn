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

func NariivchilsanUnelgee264Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "Нарийвчилсан үнэлгээ /Мэргэжилтэн / OLD",
		Identity: "id",
		Table:    "nariivchilsan_unelgee",
		Model:    new(formModels.NariivchilsanUnelgee264),
		FieldTypes: map[string]string{
			"id":                                 "Text",
			"nariin_unelgee_mn":                  "Text",
			"nariin_unelgee_en":                  "Text",
			"baiguullaga_id":                     "Select",
			"aimag_id":                           "Select",
			"sumid":                              "Select",
			"bagid":                              "Select",
			"bairshil":                           "Geographic",
			"ayul_torol_id":                      "Select",
			"ognoo":                              "Date",
			"unelgee_aanb_ner":                   "Text",
			"ui_chiglel_id":                      "Select",
			"utas":                               "Number",
			"mail":                               "Text",
			"hayg":                               "Textarea",
			"solbiltsol":                         "Geographic",
			"u_hiisen_ner_id":                    "Text",
			"u_hiisen_salbar_id":                 "Select",
			"u_hiisen_bairshil":                  "Geographic",
			"u_hiisen_argachlal_id":              "Select",
			"bb_ovog":                            "Text",
			"bb_ner":                             "Text",
			"bb_baiguullaga":                     "Text",
			"bb_alban_tushaal":                   "Text",
			"bb_register":                        "Text",
			"bb_bolovsrol":                       "Text",
			"bb_mergejil":                        "Text",
			"bb_ajil_tushilga":                   "Text",
			"bb_mail":                            "Text",
			"bb_utas":                            "Number",
			"gamshig_sudalgaa_eseh":              "Checkbox",
			"dugnelt_ur_dun":                     "Text",
			"zurag":                              "Image",
			"dugnelt_e_ognoo":                    "Date",
			"dugnelt_d_ognoo":                    "Date",
			"tolov_id":                           "Select",
			"td_avsan_ognoo":                     "Date",
			"td_materal_file":                    "File",
			"dugnelt":                            "CK",
			"tolovs_id":                          "Select",
			"tailbar":                            "CK",
			"created_at":                         "Text",
			"updated_at":                         "Text",
			"deleted_at":                         "Text",
			"sub_nariin_unelgee_dugnelt":         "SubForm",
			"sub_nariin_unelgee_file":            "SubForm",
			"unelgee_torol_id":                   "Select",
			"file_havsarsan":                     "Text",
			"sub_nariin_unelgee_bag_bureldehuun": "SubForm",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"utas": []string{},
			"mail": []string{"email"}},
		ValidationMessages: govalidator.MapData{

			"utas": []string{"mongolianMobileNumber:8 оронтой утасны дугаар оруулна уу!"},
			"mail": []string{"email:Имэйл хаягаа зөв оруулна уу!"}},
		SubForms: []map[string]interface{}{
			map[string]interface{}{
				"connection_field": "nariin_unelgee_id",
				"tableTypeColumn":  "",
				"tableTypeValue":   "",
				"table":            "sub_nariin_unelgee_dugnelt",
				"parentIdentity":   "id",
				"subIdentity":      "id",
				"subForm":          SubNariinUnelgeeDugnelt263Dataform(),
				"subFormArray":     new([]formModels.SubNariinUnelgeeDugnelt263),
			},
			map[string]interface{}{
				"connection_field": "",
				"tableTypeColumn":  "",
				"tableTypeValue":   "",
				"table":            "sub_nariin_unelgee_file",
				"parentIdentity":   "id",
				"subIdentity":      "id",
				"subForm":          SubNariinUnelgeeFileNariivchilsanUnelgee264Dataform(),
				"subFormArray":     new([]formModels.SubNariinUnelgeeFileNariivchilsanUnelgee264),
			},
			map[string]interface{}{
				"connection_field": "nariin_unelgee_id",
				"tableTypeColumn":  "",
				"tableTypeValue":   "",
				"table":            "sub_nariin_unelgee_bag_bureldehuun",
				"parentIdentity":   "id",
				"subIdentity":      "id",
				"subForm":          SubNariinUnelgeeBagBureldehuun286Dataform(),
				"subFormArray":     new([]formModels.SubNariinUnelgeeBagBureldehuun286),
			}},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}

func SubNariinUnelgeeFileNariivchilsanUnelgee264Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "Хавсралт файлууд ",
		Identity: "id",
		Table:    "sub_nariin_unelgee_file",
		Model:    new(formModels.SubNariinUnelgeeFileNariivchilsanUnelgee264),
	}
}
