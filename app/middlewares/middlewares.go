package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/helmet/v2"
)

func Set(app *fiber.App) {
	/*
		|----------------------------------------------
		| Useful MIDDLEWARES
		|----------------------------------------------
	*/
	app.Use(helmet.New())
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*, http://localhost:*, http://127.0.0.1:*",
		AllowHeaders: "Origin, Content-Type, Accept, X-Requested-With, x-csrf-token",
	}))

}
