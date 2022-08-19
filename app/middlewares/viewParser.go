package middlewares

import (
    "github.com/foolin/goview"
    "github.com/foolin/goview/supports/echoview-v4"
    "github.com/labstack/echo/v4"
    "html/template"
)

func ViewParser() echo.MiddlewareFunc {

    viewMiddleware := echoview.NewMiddleware(goview.Config{
        Root:      "views", //template root path
        Extension: ".html",
        Funcs: template.FuncMap{},
    })

    return viewMiddleware
}
