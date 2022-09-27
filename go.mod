module lambda

go 1.15

require (
	github.com/99designs/gqlgen v0.17.5
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/gofiber/fiber/v2 v2.37.0-rc.1
	github.com/gofiber/helmet/v2 v2.2.14
	github.com/gorilla/websocket v1.5.0
	github.com/lambda-platform/lambda v0.5.23
	github.com/lib/pq v1.10.4 // indirect
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/onsi/gomega v1.20.0 // indirect
	github.com/thedevsaddam/govalidator v1.9.10
	github.com/valyala/fasthttp v1.39.0
	github.com/vektah/gqlparser/v2 v2.4.2
)

//replace github.com/lambda-platform/lambda/puzzle v0.2.2 => ../../go/puzzle
//replace github.com/lambda-platform/lambda/generator v0.0.1 => ../../go/generator
//
//replace github.com/lambda-platform/lambda v0.5.23 => ../../lambda-fiber

//replace github.com/lambda-platform/arcGIS v0.0.1 => ./arcGIS
