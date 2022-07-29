package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lambda-platform/lambda/agent/agentMW"
	"lambda/app/controllers"
)

func Web(e *fiber.App) {

	//WEB ROUTE
	e.Get("/", controllers.HomeProduction) //production

	//ADMIN ROUTE
	e.Get("/admin", agentMW.IsLoggedIn(), controllers.AdminIndex(true))
}
