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

func LutSubAnimals124Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Мал амьтны дэд.төрөл",
		Identity: "id",
		Table:    "lut_sub_animals",
		Model:    new(formModels.LutSubAnimals124),
		FieldTypes: map[string]string{
			"id":          "Text",
			"animals_id":  "Select",
			"sub_animals": "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"animals_id":  []string{},
			"sub_animals": []string{}},
		ValidationMessages: govalidator.MapData{

			"animals_id":  []string{},
			"sub_animals": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
