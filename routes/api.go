package routes

import (
	"github.com/gofiber/fiber/v2"
	"lambda/app/controllers"
)

func Api(e *fiber.App) {
	a := e.Group("/api")
	a.Get("/users", controllers.Users)
	a.Post("/food-type", controllers.FoodType)
	//a.Get("/read-icons", controllers.ReadIcons)

	/* Order */
	a.Post("/create-order", controllers.CreateOrder)
	a.Get("/get-order", controllers.GetOrders)

	agent := e.Group("/auth")
	//a.Get("/", handlers.LoginPage)
	agent.Post("/login", controllers.Login)

}
