module lambda

go 1.15

require (
	github.com/gofiber/fiber/v2 v2.37.0-rc.1
	github.com/gofiber/helmet/v2 v2.2.14
	github.com/google/go-cmp v0.5.8 // indirect
	github.com/lambda-platform/lambda v0.5.32
	github.com/lib/pq v1.10.4 // indirect
	github.com/thedevsaddam/govalidator v1.9.10
	golang.org/x/net v0.0.0-20220425223048-2871e0cb64e4 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gorm.io/gorm v1.23.6
)

//replace github.com/lambda-platform/lambda/puzzle v0.2.2 => ../../go/puzzle
//replace github.com/lambda-platform/lambda/generator v0.0.1 => ../../go/generator
//
//replace github.com/lambda-platform/lambda v0.5.23 => ../../lambda-fiber

//replace github.com/lambda-platform/arcGIS v0.0.1 => ./arcGIS
