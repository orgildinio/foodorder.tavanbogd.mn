package controllers

import (
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/labstack/echo/v4"
	"net/http"
)

//home page
func HomeProduction(c echo.Context) error {
	return echoview.Render(c, http.StatusOK, "home", map[string]interface{}{})
}

func HomeData() map[string]interface{} {

	return map[string]interface{}{
		"title": "test",
	}
}
