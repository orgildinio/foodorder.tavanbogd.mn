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

func LutNervgdsenHun138Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Нэрвэгдсэн хүн амын мэдээлэл ",
		Identity: "id",
		Table:    "lut_nervgdsen_hun",
		Model:    new(formModels.LutNervgdsenHun138),
		FieldTypes: map[string]string{
			"id":             "Text",
			"nervgdsen _hun": "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"nervgdsen _hun": []string{"required"}},
		ValidationMessages: govalidator.MapData{

			"nervgdsen _hun": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
