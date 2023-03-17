package main

import (
	"github.com/lambda-platform/lambda/DBSchema"
	"lambda/bootstrap"
)

func main() {

	lambda := bootstrap.Set()
	DBSchema.GetStruct("order_detail")
	//DBSchema.GetStruct("cart_sub_menu_food")
	//DBSchema.GetStruct("sub_menu_foods_gt_neg")

	lambda.Start()

}
