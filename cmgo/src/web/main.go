package main

import (
	"fmt"

	"web/comm"
	"web/controller"
)

func init() {
	comm.InitConf()
}

func main() {
	fmt.Println(controller.Register{}.UserRegisterHandle())
	fmt.Println(comm.RedisSetKey("d","abc"))
}
