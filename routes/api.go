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
	a.Post("/create-order", agentMW.IsLoggedIn(), controllers.CreateOrder)
	a.Post("/cancel-order", agentMW.IsLoggedIn(), controllers.CancelOrder)

	//Register
	a.Post("/user-register", controllers.UserRegistration)

	//CART zahialgat
	a.Post("/cart/add", agentMW.IsLoggedIn(), controllers.AddToCart)
	a.Post("/cart/edit", agentMW.IsLoggedIn(), controllers.UpdateCart)
	a.Post("/cart/delete", agentMW.IsLoggedIn(), controllers.DeleteCart)

	//CART bagts hool
	a.Post("/cart/add-set", agentMW.IsLoggedIn(), controllers.AddCartSet)
	a.Post("/cart/edit-set", agentMW.IsLoggedIn(), controllers.UpdateCartSet)
	a.Delete("/cart/delete-set", agentMW.IsLoggedIn(), controllers.DeleteCartSet)

	//QPAY
	//a.Post("/qpay/invoice", controllers.QPayInvoice)
	//a.Post("/qpay/check", controllers.QPayPaymentCheck)
	//a.Get("/qpay/callback/:invoice_id", controllers.QPayCallBack)

	agent := e.Group("/auth")
	//a.Get("/", handlers.LoginPage)
	agent.Post("/login", controllers.Login)

}

///api/food-type
