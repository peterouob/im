package main

import (
	"gin-chat/router"
	"gin-chat/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMysql()
	utils.InitRedis()
	r := router.Router()
	//make run to start server
	r.Run(":80")
}
