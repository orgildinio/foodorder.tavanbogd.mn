module lambda

go 1.15

require (
	github.com/gofiber/fiber/v2 v2.37.0-rc.1
	github.com/gofiber/helmet/v2 v2.2.14
	github.com/lambda-platform/lambda v0.5.34
	github.com/lib/pq v1.10.4 // indirect
	github.com/thedevsaddam/govalidator v1.9.10
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gorm.io/gorm v1.24.0
)

//replace github.com/lambda-platform/lambda/puzzle v0.2.2 => ../../go/puzzle
//replace github.com/lambda-platform/lambda/generator v0.0.1 => ../../go/generator
//
replace github.com/lambda-platform/lambda v0.5.34 => ../../../lambda-fiber

//replace github.com/lambda-platform/arcGIS v0.0.1 => ./arcGIS
