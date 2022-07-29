module lambda

go 1.15

require (
	github.com/foolin/goview v0.3.0
	github.com/gofiber/fiber/v2 v2.35.0
	github.com/gofiber/helmet/v2 v2.2.14
	github.com/google/go-cmp v0.5.7 // indirect
	github.com/labstack/echo/v4 v4.7.2
	github.com/lambda-platform/lambda v0.5.21
	github.com/lib/pq v1.10.4 // indirect
	github.com/thedevsaddam/govalidator v1.9.10
)

//replace github.com/lambda-platform/lambda/puzzle v0.2.2 => ../../go/puzzle
//replace github.com/lambda-platform/lambda/generator v0.0.1 => ../../go/generator
//
//replace github.com/lambda-platform/lambda v0.5.21 => ../../lambda-fiber
