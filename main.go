package main

import (
	"lambda/bootstrap"
)

func main() {

	lambda := bootstrap.Set()
	//DBSchema.GetStruct("sub_menu_salat_gt_neg")
	//DBSchema.GetStruct("sub_menu_uuhim_gt_neg")
	//DBSchema.GetStruct("sub_menu_uuhim")
	//DBSchema.GetStruct("VB_SCHEMAS_ADMIN")
	lambda.Start()

}
