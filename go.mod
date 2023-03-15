module lambda

go 1.16

require (
	github.com/99designs/gqlgen v0.17.21
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/gofiber/fiber/v2 v2.39.0
	github.com/gofiber/helmet/v2 v2.2.18
	github.com/golang-jwt/jwt/v4 v4.4.2
	github.com/gorilla/websocket v1.5.0
	github.com/lambda-platform/lambda v0.5.93
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/onsi/gomega v1.26.0 // indirect
	github.com/thedevsaddam/govalidator v1.9.10
	github.com/valyala/fasthttp v1.40.0
	github.com/vektah/gqlparser/v2 v2.5.1
	gorm.io/gorm v1.24.5

)

//replace github.com/lambda-platform/lambda/puzzle v0.2.2 => ../../go/puzzle
//replace github.com/lambda-platform/lambda/generator v0.0.1 => ../../go/generator
//
//replace github.com/lambda-platform/lambda v0.5.50 => ../../../lambda-fiber

//replace github.com/lambda-platform/arcGIS v0.0.1 => ./arcGIS
