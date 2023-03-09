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
	a.Post("/time-order", controllers.TimeCounter)

	//Register
	a.Post("/user-register", controllers.UserRegistration)

	//CART
	a.Post("/cart/add", controllers.AddToCart)
	a.Post("/cart/edit", controllers.UpdateCart)
	a.Post("/cart/delete", controllers.DeleteCart)

	agent := e.Group("/auth")
	//a.Get("/", handlers.LoginPage)
	agent.Post("/login", controllers.Login)

}

///api/food-type
