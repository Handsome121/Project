package main

import (
	"gin-chat/initial"
	"gin-chat/router"
)

func main() {
	initial.InitConfig()
	initial.InitMysql()
	initial.InitRedis()

	r := router.Router()
	err := r.Run(":8081")
	if err != nil {
		return
	}
}
