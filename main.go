package main

import (
	"github.com/lambda-platform/lambda/DBSchema"
	"lambda/bootstrap"
)

func main() {

	lambda := bootstrap.Set()
	DBSchema.GetStruct("PASSWORD_RESETS")

	lambda.Start()

}
