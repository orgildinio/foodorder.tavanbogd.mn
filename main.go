package main

import (
	"github.com/lambda-platform/lambda/DBSchema"
	"lambda/bootstrap"
)

func main() {

	lambda := bootstrap.Set()
	DBSchema.GetStruct("view_order_detail")
	//DBSchema.GetStruct("sub_menu_uuhim_gt_neg")
	//DBSchema.GetStruct("sub_menu_uuhim")
	//DBSchema.GetStruct("VB_SCHEMAS_ADMIN")

	lambda.Start()

}
