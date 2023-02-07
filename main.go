package main

import (
    "lambda/bootstrap"
)

func main() {

    lambda := bootstrap.Set()
    //DBSchema.GetStruct("tbl_order_rule")
    //DBSchema.GetStruct("sub_menu_gt_guraw")
    //DBSchema.GetStruct("sub_menu_foods_gt_guraw")
    lambda.Start()

}
