package ebarimt

import (
	"fmt"
	"github.com/lambda-platform/ebarimt/posapi"
	"log"
)

var PosAPI *posapi.PosAPI

func init() {
	//api, err := posapi.NewPosAPI("/home/ebarimtuser/app/mmk.so")
	api, err := posapi.NewPosAPI("/home/mmk/web/ebarimt/mmk.so")
	log.Println("=======================>>>>>>>>========", api)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		PosAPI = api
		defer PosAPI.Close()
	}
}
