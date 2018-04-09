package main

import (
	"fmt"
	"web/comm"
)

func init() {
	comm.InitConf()
}

func main() {
	fmt.Println(comm.C.Db.Postresql.P_addr)
}
