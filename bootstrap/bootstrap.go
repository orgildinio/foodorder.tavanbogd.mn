package bootstrap

import (
    "github.com/labstack/echo/v4"
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

    lambdaUtils "github.com/lambda-platform/lambda/utils"
    /*
    	|----------------------------------------------
    	| App
    	|----------------------------------------------
    */
    "lambda/app/middlewares"
    "lambda/routes"
    /*
    	|----------------------------------------------
    	| Template Utils
    	|----------------------------------------------
    */
    templateUtils "github.com/lambda-platform/lambda/template/utils"

    "html/template"
)

func Set() *lambda.Lambda {

    /*
    	|----------------------------------------------
    	| Lambda Platform
    	|----------------------------------------------
    */
    lambda := lambda.New(&lambda.Settings{
        ModuleName: "lambda",
    })
    KrudMiddleWares := []echo.MiddlewareFunc{
        // arcGIS.MW(caller.GetMODEL, gridCaller.GetMODEL),
    }
    agent.Set(lambda.Echo)
    krud.Set(lambda.Echo, gridCaller.GetMODEL, caller.GetMODEL, KrudMiddleWares, true)
    exportImport.Set(lambda.Echo)
    /*
    	|----------------------------------------------
    	| MODULES
    	|----------------------------------------------
    */
    graph.Set(lambda.Echo)

    puzzle.Set(lambda.Echo, lambda.ModuleName, gridCaller.GetMODEL, false, true)
    chart.Set(lambda.Echo)
    moqup.Set(lambda.Echo)
    middlewares.Set(lambda.Echo)

    /*
    	|----------------------------------------------
    	| Template Register
    	|----------------------------------------------
    */

    templates := lambdaUtils.GetTemplates(lambda.Echo)
    TemplatePath := templateUtils.AbsolutePath()
    //* REGISTER VIEWS */
    templates["admin.html"] = template.Must(template.ParseFiles(
        TemplatePath + "views/paper.html",
    ))
    template.Must(templates["admin.html"].ParseFiles(
        "views/admin.html",
    ))

    /*
    	|----------------------------------------------
    	| ROUTES
    	|----------------------------------------------
    */
    routes.Api(lambda.Echo)
    routes.Web(lambda.Echo)

    return lambda
}
