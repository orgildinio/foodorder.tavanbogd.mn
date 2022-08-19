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

func LutMalHuis174Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "L Малын хүйс",
		Identity: "id",
		Table:    "lut_mal_huis",
		Model:    new(formModels.LutMalHuis174),
		FieldTypes: map[string]string{
			"id":       "Text",
			"mal_huis": "Text",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"mal_huis": []string{}},
		ValidationMessages: govalidator.MapData{

			"mal_huis": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
