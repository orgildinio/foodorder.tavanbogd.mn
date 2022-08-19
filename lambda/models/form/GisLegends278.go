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

func GisLegends278Dataform() dataform.Dataform {
	return dataform.Dataform{
		Name:     "Таних тэмдэг",
		Identity: "id",
		Table:    "gis_legends",
		Model:    new(formModels.GisLegends278),
		FieldTypes: map[string]string{
			"id":           "Text",
			"layer_id":     "Select",
			"element_type": "Radio",
			"style_value":  "Text",
			"style_type":   "Radio",
			"fill_color":   "ColorPicker",
			"border_color": "ColorPicker",
			"icon":         "Image",
			"line_type":    "Radio",
			"title":        "Text",
		},
		Formulas: []models.Formula{
			models.Formula{
				Targets: []models.Target{
					models.Target{
						Field: "icon",
						Prop:  "hidden",
					},
				},
				Template: "'{element_type}' == 'polygon' || '{element_type}' == 'line'",
				Model:    "",
				Form:     "main",
			}, models.Formula{
				Targets: []models.Target{
					models.Target{
						Field: "line_type",
						Prop:  "hidden",
					},
				},
				Template: "'{element_type}' == 'polygon' || '{element_type}' == 'point'",
				Model:    "",
				Form:     "main",
			}, models.Formula{
				Targets: []models.Target{
					models.Target{
						Field: "style_type",
						Prop:  "hidden",
					}, models.Target{
						Field: "border_color",
						Prop:  "hidden",
					},
				},
				Template: "'{element_type}' == 'line' || '{element_type}' == 'point'",
				Model:    "",
				Form:     "main",
			}, models.Formula{
				Targets: []models.Target{
					models.Target{
						Field: "fill_color",
						Prop:  "hidden",
					},
				},
				Template: "'{element_type}' == 'point'",
				Model:    "",
				Form:     "main",
			},
		},
		ValidationRules: govalidator.MapData{

			"layer_id":     []string{},
			"fill_color":   []string{},
			"border_color": []string{},
			"icon":         []string{}},
		ValidationMessages: govalidator.MapData{

			"layer_id":     []string{},
			"fill_color":   []string{},
			"border_color": []string{},
			"icon":         []string{}},
		SubForms:         []map[string]interface{}{},
		AfterInsert:      nil,
		AfterUpdate:      nil,
		BeforeInsert:     nil,
		BeforeUpdate:     nil,
		TriggerNameSpace: "",
	}
}
