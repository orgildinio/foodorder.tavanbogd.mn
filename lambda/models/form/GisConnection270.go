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

func GisConnection270Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "Газрын зургийн сервис холболт",
		Identity: "id",
		Table:    "gis_connection",
		Model:    new(formModels.GisConnection270),
		FieldTypes: map[string]string{
			"id":         "Text",
			"title":      "Text",
			"connection": "arcGISLayerAttributeConnector",
			"layer":      "arcGISLayerAttributeConnectorLayer",
			"local_form": "Select",
			"local_grid": "Select",
		},
		Formulas: []models.Formula{},
		ValidationRules: govalidator.MapData{

			"title":      []string{},
			"connection": []string{},
			"layer":      []string{},
			"local_form": []string{},
			"local_grid": []string{}},
		ValidationMessages: govalidator.MapData{

			"title":      []string{},
			"connection": []string{},
			"layer":      []string{},
			"local_form": []string{},
			"local_grid": []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
