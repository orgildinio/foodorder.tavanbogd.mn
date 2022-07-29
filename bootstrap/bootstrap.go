package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lambda-platform/lambda/chart"
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
		| Graphql
		|----------------------------------------------
	*/
	"lambda/lambda/graph"
	/*
		|----------------------------------------------
		| PRO MODULES
		|----------------------------------------------
	*/

	"github.com/lambda-platform/lambda/moqup"

	/*
		|----------------------------------------------
		| App
		|----------------------------------------------
	*/
	"lambda/app/middlewares"
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
	KrudMiddleWares := []fiber.Handler{
		// arcGIS.MW(caller.GetMODEL, gridCaller.GetMODEL),
	}
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
	middlewares.Set(Lambda.App)
	/*
		|----------------------------------------------
		| ROUTES
		|----------------------------------------------
	*/
	routes.Api(Lambda.App)
	routes.Web(Lambda.App)

	return Lambda
}
