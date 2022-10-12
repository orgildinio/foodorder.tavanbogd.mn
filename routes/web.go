package routes

import (
	"github.com/gofiber/fiber/v2"
	"lambda/app/controllers"
)

func Web(e *fiber.App) {

	//WEB ROUTE
	e.Get("/", controllers.Home)                  //production
	e.Get("/admin/p/:menu_id", controllers.Admin) //production
	e.Get("/auth/forgot", controllers.Login)
	e.Get("/auth/login", controllers.Login)
	//ADMIN ROUTE
	//e.Get("/admin", agentMW.IsLoggedIn(), controllers.AdminIndex(true))
}
