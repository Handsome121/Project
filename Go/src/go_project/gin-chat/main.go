package main

import (
	"gin-chat/initial"
	"gin-chat/router"
)

func main() {
	initial.InitConfig()
	initial.InitMysql()

	r := router.Router()
	err := r.Run()
	if err != nil {
		return
	}
}