package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lambda-platform/lambda/dataform"
	"github.com/lambda-platform/lambda/datagrid"
	"lambda/app/controllers"
)

type crudResponse struct {
	Data struct {
		ID int `gorm:"column:id;" json:"id"`
	} `json:"data"`
}

func ReportCalculator(c *fiber.Ctx, GridModelCaller func(schema_id string) datagrid.Datagrid, FormModelCaller func(schema_id string) dataform.Dataform) error {

	if c.Path() == "/lambda/puzzle/grid/data/313" {
		grid := GridModelCaller("313")
		return controllers.HuraanguiTailanFetchData(c, grid)
	} else if c.Path() == "/lambda/puzzle/grid/aggergation/313" {
		grid := GridModelCaller("313")
		return controllers.HuraanguiTailanAggergation(c, grid)
	}

	return c.Next()
}

func ReportCalculatorMW(GetGridMODEL func(schema_id string) datagrid.Datagrid, GetMODEL func(schema_id string) dataform.Dataform) fiber.Handler {

	return func(c *fiber.Ctx) error {
		return ReportCalculator(c, GetGridMODEL, GetMODEL)
	}

}
