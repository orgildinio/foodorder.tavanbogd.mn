package routes

import (
	"github.com/gofiber/fiber/v2"
	"lambda/app/controllers"
)

func Web(e *fiber.App) {

	//WEB ROUTE
	e.Get("/", controllers.Home)                  //production
	e.Get("/admin/p/:menu_id", controllers.Admin) //production

	//ADMIN ROUTE
	//e.Get("/admin", agentMW.IsLoggedIn(), controllers.AdminIndex(true))

	e.Get("/auth/forgot", controllers.LoginPage)
	e.Get("/auth/login", controllers.LoginPage)
}
