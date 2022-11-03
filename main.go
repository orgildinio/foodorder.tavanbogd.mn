package main

import (
	"lambda/bootstrap"
)

func main() {

	lambda := bootstrap.Set()
	//DBSchema.GetStruct("USERS")
	//DBSchema.GetStruct("VB_SCHEMAS_ADMIN")
	lambda.Start()

}
