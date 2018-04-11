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
	fmt.Println(controller.Register{}.UserRegisterHandle()) 	//插入mongodb
	fmt.Println(comm.RedisSetKey("caimin", "彩民"))   //插入redis
	fmt.Println(comm.RedisGetKey("d"))                      //读取redis

	comm.RedisSetMap()

	i := comm.RedisAccumulation("kk")
	fmt.Println(i)
}
