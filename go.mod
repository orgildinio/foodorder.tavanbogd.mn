module lambda

go 1.16

require (
	github.com/gofiber/fiber/v2 v2.39.0
	github.com/gofiber/helmet/v2 v2.2.18
	github.com/golang-jwt/jwt/v4 v4.4.2
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/lambda-platform/lambda v0.5.79
	github.com/thedevsaddam/govalidator v1.9.10
	golang.org/x/net v0.2.0 // indirect
	golang.org/x/tools v0.2.0 // indirect
	gorm.io/gorm v1.24.0

)

//replace github.com/lambda-platform/lambda/puzzle v0.2.2 => ../../go/puzzle
//replace github.com/lambda-platform/lambda/generator v0.0.1 => ../../go/generator
//
//replace github.com/lambda-platform/lambda v0.5.50 => ../../../lambda-fiber

//replace github.com/lambda-platform/arcGIS v0.0.1 => ./arcGIS
