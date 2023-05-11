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
	a.Get("/send-notification", controllers.SendNotification)
	//a.Get("/read-icons", controllers.ReadIcons)

	//Register
	a.Post("/user-register", controllers.UserRegistration)

	//CART zahialgat

	a.Post("/cart/add", agentMW.IsLoggedIn(), controllers.AddToCart)
	a.Post("/cart/edit", agentMW.IsLoggedIn(), controllers.UpdateCart)
	a.Post("/cart/delete", agentMW.IsLoggedIn(), controllers.DeleteCart)

	/* Cart Set */
	a.Post("/cart/create-set", agentMW.IsLoggedIn(), controllers.AddToCartSet)
	a.Post("/cart/edit-set", agentMW.IsLoggedIn(), controllers.EditCartItem)
	a.Post("/cart/delete-set", agentMW.IsLoggedIn(), controllers.DeleteCartItem)

	/* Order Zahialgat */
	a.Get("/pending/order", agentMW.IsLoggedIn(), controllers.GetPendingOrder)
	a.Get("/create-order", agentMW.IsLoggedIn(), controllers.CreateOrder)
	a.Post("/cancel-order", agentMW.IsLoggedIn(), controllers.CancelOrder)

	/* Order Zahialgat */
	//a.Get("/create-order-set", agentMW.IsLoggedIn(), controllers.CreateOrderSet)

	//QPAY
	a.Post("/qpay/invoice", controllers.QPayInvoice)
	//a.Post("/qpay/check", controllers.QPayPaymentCheck)
	a.Get("/qpay/callback/:invoice_id", agentMW.IsLoggedIn(), controllers.QPayCallBack)

	//Latest Payment
	a.Get("/payment/latest-payment", agentMW.IsLoggedIn(), controllers.LaterPay)

	//Reception
	a.Post("/reception", agentMW.IsLoggedIn(), controllers.RecepcionRequest)

	//CART bagts hool
	/* NOT USED */
	//a.Post("/cart/add-set", agentMW.IsLoggedIn(), controllers.AddCartSet)
	//a.Post("/cart/edit-set", agentMW.IsLoggedIn(), controllers.UpdateCartSet)
	//a.Delete("/cart/delete-set", agentMW.IsLoggedIn(), controllers.DeleteCartSet)
	/* NOT USED */

	agent := e.Group("/auth")
	//a.Get("/", handlers.LoginPage)
	agent.Post("/login", controllers.Login)

	//Ebarimt

	a.Get("/ebarimt-info", controllers.EBarimtInfo)
	a.Get("/ebarimt-check", controllers.EBarimtCheck)
	a.Get("/ebarimt-send", controllers.EBarimtSend)
	a.Get("/ebarimt-put", controllers.EBarimtPut)

}

///api/food-type
