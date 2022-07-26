module lambda

go 1.15

require (
	github.com/foolin/goview v0.3.0
	github.com/google/go-cmp v0.5.7 // indirect
	github.com/labstack/echo/v4 v4.7.2
	github.com/lambda-platform/lambda v0.5.19
	github.com/lib/pq v1.10.4 // indirect
	github.com/thedevsaddam/govalidator v1.9.10
	golang.org/x/net v0.0.0-20220225172249-27dd8689420f // indirect
	golang.org/x/sys v0.0.0-20220111092808-5a964db01320 // indirect
)

//replace github.com/lambda-platform/lambda/puzzle v0.2.2 => ../../go/puzzle
//replace github.com/lambda-platform/lambda/generator v0.0.1 => ../../go/generator
//
//replace github.com/lambda-platform/lambda v0.5.8 => ../../lambda-go
