package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lambda-platform/lambda/agent/agentMW"
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
	a.Post("/cart/add", agentMW.IsLoggedIn(), controllers.AddToCart)
	a.Post("/cart/edit", agentMW.IsLoggedIn(), controllers.UpdateCart)
	a.Post("/cart/delete", agentMW.IsLoggedIn(), controllers.DeleteCart)

	a.Post("/cart/add-set", controllers.AddCartSet)
	a.Post("/cart/edit-set", controllers.UpdateCartSet)
	a.Delete("/cart/delete-set", controllers.DeleteCartSet)

	agent := e.Group("/auth")
	//a.Get("/", handlers.LoginPage)
	agent.Post("/login", controllers.Login)

}

///api/food-type
