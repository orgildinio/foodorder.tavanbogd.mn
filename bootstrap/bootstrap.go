package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lambda-platform/lambda/chart"
	"lambda/app/middlewares"
	"lambda/lambda/graph"
	"lambda/lambda/models/form/caller"
	gridCaller "lambda/lambda/models/grid/caller"
	/*
		|----------------------------------------------
		| Generated Models
		|----------------------------------------------
	*/
	"github.com/lambda-platform/lambda"
	"github.com/lambda-platform/lambda/agent"
	"github.com/lambda-platform/lambda/exportImport"
	"github.com/lambda-platform/lambda/krud"
	"github.com/lambda-platform/lambda/puzzle"

	/*
		|----------------------------------------------
		| PRO MODULES
		|----------------------------------------------
	*/

	"github.com/lambda-platform/lambda/moqup"
	"lambda/routes"
)

func Set() *lambda.Lambda {
	/*
		|----------------------------------------------
		| Lambda Platform
		|----------------------------------------------
	*/
	Lambda := lambda.New(&lambda.Settings{
		ModuleName: "lambda",
	})
	middlewares.Set(Lambda.App)
	KrudMiddleWares := []fiber.Handler{}
	routes.Api(Lambda.App)
	agent.Set(Lambda.App)
	krud.Set(Lambda.App, gridCaller.GetMODEL, caller.GetMODEL, KrudMiddleWares, true)

	exportImport.Set(Lambda.App)
	/*
		|----------------------------------------------
		| MODULES
		|----------------------------------------------
	*/
	graph.Set(Lambda.App)
	puzzle.Set(Lambda.App, Lambda.ModuleName, gridCaller.GetMODEL, false, true)
	chart.Set(Lambda.App)
	moqup.Set(Lambda.App)

	/*
		|----------------------------------------------
		| ROUTES
		|----------------------------------------------
	*/

	routes.Web(Lambda.App)

	return Lambda
}
